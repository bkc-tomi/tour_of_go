# goroutine と Channel
エクササイズではありませんが、難しく感じたのでまとめようと思います。


## goroutine
goroutineはgoで並行処理を行う仕組みのことである。<br>
並行処理とは、複数の異なった処理を同時に実行することである。以下はgoroutineを使った例である。
```
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("w")
	go say("o")
	go say("r")
	go say("l")
	go say("d")
	say("end")
}
/*
実行結果
d
end
o
r
w
l
end
d
r
l
o
w
d
l
r
w
o
end
l
w
end
r
o
d
d
l
w
o
r
end
*/
```
このように並行処理goroutineを使いたい場合は関数名などの前にgoをつけるだけである。1秒刻みで5回引数に与えられた単語を表示する関数を「w, o, r, l, d」のそれぞれの文字で並行処理を行い、その後並行処理とは別にendに対して行っている。<br>
実行結果をみたらわかるように、同時に処理を行い、処理が終わったものから表示しているのでバラバラに表示されている。

後から触れると思うが、非同期処理も可能である。



## channel
Channel とは

- Channel は goroutine 間でのメッセージパッシングをするためのもの<br>
個人的見解
「関数はスコープが閉じているので変数を受け渡せない。
goroutineは並行処理なのでどの処理から先に完了するか分からない。
といった点から値のやりとりには特殊な変数が必要。それがchannel」

- メッセージの型を指定できる
- first class value であり、引数や戻り値にも使える
- send/receive でブロックする
- buffer で、一度に扱えるメッセージ量を指定できる

channelはmake(chan 型)で宣言し <- で受信・送信を行う。<br>
ch <- :受信<br>
<- ch :送信<br>
以下のコードと実行結果を見れば、(1)でch自身を出力しているのを見れば、チャンネルはポインター的な役割なのがわかる。
非同期処理内で指定した型の値を受け取る。処理の外で別の変数に取得した値を渡す。

```
func main() {
	ch := make(chan int)
	go func() { ch <- 1 }()
	fmt.Println(ch) // (1)
	v := <-ch
	fmt.Println(v) // (2)
}

/*
実行結果
:~/go/src/tour_of_go/channel/ch1 (master)
$ ./ch1
0xc00008c060
1
*/
```
以下では２つのgoroutineに対して２つのチャンネルを定義している。goroutineは並行処理なので実行まいに結果表示の順番が異なる。
```
func task1(v string, ch chan int) {
	fmt.Println(v)
	ch <- 1
}

func task2(v string, ch chan int) {
	fmt.Println(v)
	ch <- 2
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go task2("task2", ch2)
	go task1("task1", ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
}

/*
実行結果
:~/go/src/tour_of_go/channel/ch3 (master)
$ ./ch3
task1
task2
1
2
:~/go/src/tour_of_go/channel/ch3 (master)
$ ./ch3
task1
1
task2
2
*/
```
bufferはchannelが同時に受診できる値の大きさを示す。下の例でいくと２つ。<br>
ここでは、３つの値を送信を挟まずに受信してしまっているので同時に３つの値を保持することになりエラーになる。
```
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
/*
実行結果
:~/go/src/tour_of_go/channel/buffer (master)
$ ./buffer 
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        /Users/go/src/tour_of_go/channel/buffer/buffer.go:11 +0x9b
*/
```
同時に持てる量なので以下のようにchannelからそのつど送信を行ったあとだと受信することが出来る。
```
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
実行結果
1
2
3
```

## range and close
close(ch chan)でチャンネルを閉じることが出来る。チャンネルを閉じると値を受信することが出来なくなる。チャンネルを閉じた後も値の送信は出来る。<br>
チャンネルからは値と真偽値を受け取ることが出来ます。v, ok := <-ch
```
func a(c chan int) {
	c <- 1
}

func main() {
	ch := make(chan int, 10)
	a(ch)
	close(ch)
	v, ok := <-ch

	if ok {
		fmt.Println(v, ok)
	} else {
		fmt.Println("channel closed.")
	}
}
/*
実行結果
1 true
*/
```
```
func a(c chan int) {
	c <- 1
}

func main() {
	ch := make(chan int, 10)
	a(ch)
	close(ch)
	a(ch)
	v, ok := <-ch

	if ok {
		fmt.Println(v, ok)
	} else {
		fmt.Println("channel closed.")
	}
}
/*
実行結果
panic: send on closed channel

goroutine 1 [running]:
main.a(...)
	/tmp/sandbox643135727/prog.go:8
main.main()
	/tmp/sandbox643135727/prog.go:20 +0x91
*/
```
チャンネルを閉じるのは、呼び出す先の関数に記述していても機能する。
```
func b(c chan int) {
	c <- 2
	close(c)
}

func main() {
	ch := make(chan int, 10)
	b(ch)
	b(ch)
	v, ok := <-ch

	if ok {
		fmt.Println(v, ok)
	} else {
		fmt.Println("channel closed.")
	}
}
/*
実行結果
panic: send on closed channel

goroutine 1 [running]:
main.b(...)
	/tmp/sandbox201782908/prog.go:12
main.main()
	/tmp/sandbox201782908/prog.go:19 +0x91
*/
```
何も値を受信しておらず、かつチャンネルが閉じている時はv, ok := <-chのokはfalseを受信する。
```
func main() {
	ch := make(chan int, 10)
	close(ch)
	v, ok := <-ch

	fmt.Println(v, ok)
}
/*
実行結果
0 false
*/
```
複数の値を受信している場合はrangeでループを回すことが出来る。この時ループはチャンネルが閉じられているところまで繰り返される。なのでrangeで回す時は、その前にclose()で必ず閉じておく必要がある。閉じていなかった場合は、パニックを起こします。
```
func c(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
}

func main() {
	ch := make(chan int, 10)
	c(ch)
	close(ch)
	for c := range ch {
		fmt.Println(c)
	}
}
/*
実行結果
0
1
2
3
4
5
6
7
8
9
*/
```
注：チャネルを閉じる必要があるのは送信側だけで、受信側は閉じないでください。閉じたチャネルで送信すると、パニックが発生します。
<br>
別のメモ：チャネルはファイルとは異なります。通常は閉じる必要はありません。閉じる必要があるのは、rangeループを終了するなど、受信側に値が来ないことを通知する必要がある場合のみです。