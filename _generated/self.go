package _generated

//go:generate zebrapack

type ZError struct {
	Context int `zid:"0"`
}

type ZResponse struct {
	Error ZError `zid:"0"`
}
