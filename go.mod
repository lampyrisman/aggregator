module main

go 1.18

replace skv => ./modules/skv

replace structs => ./modules/structs

require (
	skv v0.0.0-00010101000000-000000000000
	structs v0.0.0-00010101000000-000000000000
)
