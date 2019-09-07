package testpkg


type SrcStruct struct {
	p1 string
}

func NewSrcStruct(p1 string) *SrcStruct {
	return &SrcStruct{p1: p1}
}

