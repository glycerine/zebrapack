package testdata

import (
	"time"
)

//go:generate zebrapack

// interface testing
type Familiar interface {
	Name() string
}

type Dog struct{}

func (d *Dog) Name() string { return "Rover" }

type Owl struct{}

func (o *Owl) Name() string { return "Hedwig" }

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
	Pet Familiar  `zid:"4"`
}
