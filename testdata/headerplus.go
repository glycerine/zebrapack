package testdata

import (
	"fmt"
	"time"

	"github.com/glycerine/zebrapack/msgp"
)

//go:generate zebrapack

// interface testing
type Familiar interface {
	Name() string

	msgp.Decodable
	msgp.Encodable
	msgp.Marshaler
	msgp.Unmarshaler
	msgp.Sizer
}

type Dog struct {
	Nm string `zid:"0"`
}

func (d *Dog) Name() string { return d.Nm }

type Owl struct {
	Nm string `zid:"0"`
}

func (o *Owl) Name() string { return o.Nm }

type Header struct {
	S1 string `zid:"0"`
	S2 string `zid:"1"`
	C  int64  `zid:"2"`
}

type Plus struct {
	Tm  time.Time `zid:"0"`
	N   int64     `zid:"1"`
	S   string    `zid:"2"`
	F   float64   `zid:"3"`
	Pet Familiar  `msg:",iface" zid:"4"`
}

func (pl *Plus) NewValueAsInterface(typename string) interface{} {
	fmt.Printf("\n DEBUG! NewValueAsInterface called with typename='%s'\n", typename)
	switch typename {
	case "Dog":
		return &Dog{}
	case "Owl":
		return &Owl{}
	}
	panic(fmt.Sprintf("unknown "))
}
