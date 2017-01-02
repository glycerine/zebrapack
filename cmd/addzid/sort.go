package main

type ByFinalOrder []*Field

func (s ByFinalOrder) Len() int {
	return len(s)
}
func (s ByFinalOrder) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByFinalOrder) Less(i, j int) bool {
	return s[i].finalOrder < s[j].finalOrder
}

type ByOrderOfAppearance []*Field

func (s ByOrderOfAppearance) Len() int {
	return len(s)
}
func (s ByOrderOfAppearance) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByOrderOfAppearance) Less(i, j int) bool {
	return s[i].orderOfAppearance < s[j].orderOfAppearance
}

type ByGoName []*Struct

func (s ByGoName) Len() int {
	return len(s)
}
func (s ByGoName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByGoName) Less(i, j int) bool {
	return s[i].goName < s[j].goName
}

type AlphaHelper struct {
	Name string
	Code []byte
}

type AlphaHelperSlice []AlphaHelper

func (s AlphaHelperSlice) Len() int {
	return len(s)
}
func (s AlphaHelperSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s AlphaHelperSlice) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}
