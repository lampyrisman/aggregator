package structs

// avail actions: get set resp err ping

type KV struct {
	Ip       string
	Port     string
	LastSeen int32
}

type Reg_data struct {
	Action    string `json:"action"`
	Role      string `json:"role,omitempty"`
	IP        string `json:"ip,omitempty"`
	Port      string `json:"port,omitempty"`
	Timestamp int32
	Result    string
}
