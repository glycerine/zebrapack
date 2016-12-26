package main

import (
	"time"
)

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
	Slice     []S2        `zid:"0"`
	Transform map[int]*S2 `zid:"1"`
	Myptr     *S2         `zid:"2"`
	Myarray   [3]string   `zid:"3"`
	MySlice   []string    `zid:"4"`
}

type A struct {
	Name   string    `zid:"0"`
	Bday   time.Time `zid:"1"`
	Phone  string    `zid:"2"`
	Sibs   int       `zid:"3"`
	GPA    float64   `zid:"4"`
	Friend bool      `zid:"5"`
}
