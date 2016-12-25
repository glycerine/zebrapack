package main

const zebraSchemaId64 = 0x6eb25cc0f9a3e

func main() {}

type S struct{}

type S2 struct {
	A struct{}         `zid:"0" msg:"alpha"`
	B string           `msg:"beta" zid:"1"`
	R map[string]uint8 `zid:"2" msg:"ralph"`
	P uint16           `zid:"3" deprecated:"true"`
	Q uint32           `zid:"4"`
	T int64            `zid:"5"`
}

type Big struct {
	Slice []S2 `zid:"0"`
}
