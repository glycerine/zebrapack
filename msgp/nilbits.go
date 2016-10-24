package msgp

import (
	"fmt"
	//"runtime/debug"
)

const MaxNestedStructPointerDepth = 8

// NilBitsStack is a helper for Unmarshal
// methods to track where we are when
// deserializing from empty/nil/missing
// fields.
type NilBitsStack struct {
	// simulate getting nils on the wire
	AlwaysNil     bool
	LifoAlwaysNil int                                 // this lifo always had false in it, so change it to a count.
	LifoBts       [MaxNestedStructPointerDepth][]byte // caps our static tree depth at 8, but avoids all memory allocation during the decode hotpath.

	// UnsafeZeroCopy will make strings point to the
	// original msgpack buffers. This is unsafe because
	// if the buffer changes the string will change/be
	// invalid/not protected against re-use. But for
	// processing and disposing of single messages, one-at-at-time,
	// without re-using any part of a message (or making a copy of strings explicitly with copy()
	// if you must) then we can avoid all allocations for strings.
	UnsafeZeroCopy bool
}

func (r *NilBitsStack) Init(cfg *RuntimeConfig) {
	if cfg != nil {
		r.UnsafeZeroCopy = cfg.UnsafeZeroCopy
	}
}

func (r *NilBitsStack) IsNil(bts []byte) bool {
	if r.AlwaysNil {
		return true
	}
	if len(bts) != 0 && bts[0] == mnil {
		return true
	}
	return false
}

// OnlyNilSlice is a slice that contains
// only the msgpack nil (0xc0) bytes.
var OnlyNilSlice = []byte{mnil}

// AlwaysNilString returns a string representation
// of the internal state of the NilBitsStack for
// debugging purposes.
func (r *NilBitsStack) AlwaysNilString() string {

	s := "bottom: "
	for i := 0; i > r.LifoAlwaysNil; i++ {
		s += "T"
	}
	/*	for _, v := range r.LifoAlwaysNil {
			if v {
				s += "T"
			} else {
				s += "f"
			}
		}
	*/
	return s
}

// PushAlwaysNil will set r.AlwaysNil to true
// and store bts on the internal stack.
func (r *NilBitsStack) PushAlwaysNil(bts []byte) []byte {
	//fmt.Printf("PushAlwaysNil(), pre we are '%v'. Called from: stack is '%v'\n", r.AlwaysNilString(), string(debug.Stack()))
	// save current state

	if r.LifoAlwaysNil == MaxNestedStructPointerDepth {
		panic(fmt.Sprintf("we hit our maximum nested struct-inside-struct depth! you must recompile msgp with github.com/glycerine/zebrapack/msgp/nilbits.go: MaxNestedStructPointerDepth set to greater than the current value of %v, and then regenerate your msgp Unmarshaling code.", MaxNestedStructPointerDepth))
	}

	r.LifoBts[r.LifoAlwaysNil] = bts
	r.LifoAlwaysNil++

	// set reader r to always return nils
	r.AlwaysNil = true

	return OnlyNilSlice
}

// PopAlwaysNil pops the last []byte off the internal
// stack and returns it. If the stack is empty
// we panic.
func (r *NilBitsStack) PopAlwaysNil() (bts []byte) {
	//fmt.Printf("my NilBitsTack.PopAlwaysNil() called!! ...  stack is '%v'\n", string(debug.Stack()))
	//	defer func() {
	//		fmt.Printf("len of bts returned by PopAlwaysNil: %v, and debug string: '%v'\n",
	//			len(bts), r.AlwaysNilString())
	//	}()
	n := r.LifoAlwaysNil
	if n == 0 {
		panic("PopAlwaysNil called on empty lifo")
	}
	if n < 0 {
		panic("PopAlwaysNil called on illegal-less-thab-empty lifo")
	}
	bts = r.LifoBts[n-1]

	r.LifoAlwaysNil--
	if r.LifoAlwaysNil == 0 {
		r.AlwaysNil = false
	}

	return bts
}

func ShowFound(found []bool) string {
	s := "["
	for i := range found {
		if found[i] {
			s += "1,"
		} else {
			s += "0,"
		}
	}
	s += "]"
	return s
}
