package main

const zebraSchemaId64 = 0xfffeeeeddd

func main() {}

type S struct{}

type S2 struct {
	A struct{}         `zid:"0"`
	B string           `zid:"1"`
	R map[string]uint8 `zid:"2"`
	P uint16           `zid:"3"`
	Q uint32           `zid:"4"`
	T int64            `zid:"5"`
}

type Big struct {
	Slice []S2 `zid:"0"`
}
