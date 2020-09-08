package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (a *rot13Reader) Read(rb []byte) (int, error) {
	n, e := a.r.Read(rb)
	for i, v := range rb {
		if v >= 'A' && v <= 'Z' {
			rb[i] = (v-'A'+13)%26 + 'A'
		}
		if v >= 'a' && v <= 'z' {
			rb[i] = (v-'a'+13)%26 + 'a'
		}
	}
	return n, e
}

func main() {
	// ここのファイルストリームの操作の理解が不十分
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
