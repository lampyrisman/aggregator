package skv

import (
	"fmt"
	s "structs"
)

func cerr(err error) (isErr bool) {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

func setData(from *s.Reg_data, servers map[string]s.KV) {
	servers[from.IP+":"+from.Port] = s.KV{
		Ip:       from.IP,
		Port:     from.Port,
		LastSeen: 0,
	}
	from.Action = "RESPONSE"
	from.Result = "OK"
}

func prestart(c_from chan s.Reg_data, servers map[string]s.KV) {
	var resp s.Reg_data
	from := <-c_from
	if from.Action == "set" {
		setData(&from, servers)
	} else {
		fmt.Println("No one server in online")
		resp.Action = "ERR"
		resp.Result = "NOSRV"
		c_from <- resp
		return
	}
	resp.Action = "RESPONSE"
	resp.Result = "OK"

	c_from <- resp
}

func Reg(c chan s.Reg_data) {
	fmt.Println("Start")
	servers := make(map[string]s.KV)
	for {

		if len(servers) == 0 {
			prestart(c, servers)
			continue
		}

		for _, v := range servers {
			rcv_data := <-c
			if rcv_data.Action == "set" {
				setData(&rcv_data, servers)
			}
			if rcv_data.Action == "get" {
				rcv_data = s.Reg_data{}
				rcv_data.Action = "RESPONSE"
				rcv_data.IP = v.Ip
				rcv_data.Port = v.Port
				rcv_data.Result = "OK"
			}

			fmt.Print("skv got", rcv_data, "\n")
			//			fmt.Println("KV:", servers, v)
			c <- rcv_data
		}
	}
}
