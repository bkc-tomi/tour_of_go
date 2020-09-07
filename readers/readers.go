package main

import (
	"golang.org/x/tour/reader"
)

// MyReader type
type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(rb []byte) (int, error) {
	var n int
	var e error
	for n, e = 0, nil; n < len(rb); n++ {
		// 配列はミュータブルなのでポインタを使用せず要素の書き換えが可能
		rb[n] = 'A'
	}
	return n, e
}

func main() {
	reader.Validate(MyReader{})
}
