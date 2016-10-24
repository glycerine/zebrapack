package msgp

type RuntimeConfig struct {

	// UnsafeZeroCopy will make strings point to the
	// original msgpack buffers. This is unsafe because
	// if the buffer changes the string will change/be
	// invalid/not protected against re-use. But for
	// processing and disposing of single messages, one-at-at-time,
	// without re-using any part of a message (or making a copy of strings explicitly with copy()
	// if you must) then we can avoid all allocations for strings.
	UnsafeZeroCopy bool
}
