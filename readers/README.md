# Readers
またまたメソッド・インターフェースのエクササイズです。今回はファイル操作に使われるっぽいio.Readerインターフェースを使ったエクササイズです。これはbyte列を読み出すRead()メソッドを持つことを保証します。

今回はこのメソッドを使用して与えられたbyte列分のASCII文字コードの'A'を排出するMyReader型を作成します。
main()関数内で"golang.org/x/tour/reader"のパッケージを使用して実装したMyReader型がちゃんと実装されているかをチェックします。以下は元のコードです。見てみると1KB分のbyte列を１MB分1024回チェックしてきちんと実装されているかを確認しています。
reader.Validate()
```
func Validate(r io.Reader) {
	b := make([]byte, 1024, 2048)
	i, o := 0, 0
	for ; i < 1<<20 && o < 1<<20; i++ { // test 1mb
		n, err := r.Read(b)
		fmt.Println(i, o, n)
		for i, v := range b[:n] {
			if v != 'A' {
				fmt.Fprintf(os.Stderr, "got byte %x at offset %v, want 'A'\n", v, o+i)
				return
			}
		}
		o += n
		if err != nil {
			fmt.Fprintf(os.Stderr, "read error: %v\n", err)
			return
		}
	}
	if o == 0 {
		fmt.Fprintf(os.Stderr, "read zero bytes after %d Read calls\n", i)
		return
	}
	fmt.Println("OK!")
}
```

[Exercise: Readers](https://tour.golang.org/methods/22)