package mapto

import (
	"github.com/juntaki/mapto/testpkg"
	"github.com/juntaki/pp"
	"reflect"
	"testing"
)

func TestMapBasic(t *testing.T) {
	src := testpkg.NewSrcStruct("p1")
	type destStruct struct {
		P1 string `mapto:"p1"`
		P2 int    `mapto:"p2"`
	}
	dest := &destStruct{}
	want := &destStruct{
		P1: "p1",
		P2: 0,
	}
	Map(dest, src)
	if !reflect.DeepEqual(dest, want) {
		pp.Println(dest, want)
		t.Fatal("fail to get expected output")
	}
}

func TestMapPointer(t *testing.T) {
	type srcStruct struct {
		p1 *string
	}
	p1 := "p1"
	src := &srcStruct{
		p1: &p1,
	}
	type destStruct struct {
		P1 *string `mapto:"p1"`
	}
	dest := &destStruct{}
	want := &destStruct{
		P1: &p1,
	}
	Map(dest, src)
	if !reflect.DeepEqual(dest, want) {
		pp.Println(dest, want)
		t.Fatal("fail to get expected output")
	}
}

func TestMapNested(t *testing.T) {
	type srcStructInner struct {
		p1 string
		p2 int
	}
	type srcStruct struct {
		p1 srcStructInner
		p2 *srcStructInner
	}
	src := &srcStruct{
		p1: srcStructInner{
			p1: "p1p1",
			p2: 11,
		},
		p2: &srcStructInner{
			p1: "p2p1",
			p2: 22,
		},
	}
	type destStruct struct {
		P1 string `mapto:"p1.p1"`
		P2 int    `mapto:"p2.p2"`
	}
	dest := &destStruct{}
	want := &destStruct{
		P1: "p1p1",
		P2: 22,
	}
	Map(dest, src)
	if !reflect.DeepEqual(dest, want) {
		pp.Println(dest, want)
		t.Fatal("fail to get expected output")
	}
}
