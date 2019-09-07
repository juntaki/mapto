# mapto

You may want to keep the structure field private, but you must map the value to the target structure.
Making a lot of getters is one of the solutions, but you can keep the source code short with a few hacks.

# How to use

https://play.golang.org/p/BSQJZQ8U0Ya

```
go get github.com/juntaki/mapto
```

Source struct (with private fields)

```
type SrcStruct struct {
	p1 string
}
```

target (with "mapto" tag)

```
type destStruct struct {
	P1 string `mapto:"p1"`
}
```

Do the hack

```
src := testpkg.NewSrcStruct("p1")
dest := &destStruct{}

mapto.Map(dest, src)

fmt.Println(dest.P1) // "p1"
```