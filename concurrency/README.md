# goroutine と Channel
エクササイズではありませんが、難しく感じたのでまとめようと思います。


## goroutine(ゴルーチン)
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
略
w
o
r
end
*/
```
このように並行処理goroutineを使いたい場合は関数名などの前にgoをつけるだけである。1秒刻みで5回引数に与えられた単語を表示する関数を「w, o, r, l, d」のそれぞれの文字で並行処理を行い、その後並行処理とは別にendに対して行っている。<br>
実行結果をみたらわかるように、同時に処理を行い、処理が終わったものから表示しているのでバラバラに表示されている。


## channel
Channel とは

Channel は goroutine 間でのメッセージパッシングをするためのもの<br>

基本的にコンピュータはコードの上から書かれたものから実行していく。前の処理が完了しないと次の処理は実行されない。goroutineは関数の前にgoがついているものを同時に実行する。そうした場合、二つの異なるgoroutineが一つの値を書き換える場合問題が起きる。

<br>

例えば以下のように、与えられた引数をインクリメントするinc関数を実装し、goroutineで1000回実行する。
比較のためにそれをさらにループで5回繰り返し、5回の処理結果を比べてみる。
```
func inc(x *int) {
	*x++
}

func main() {
	var x int
	for h := 0; h < 5; h++ {
		x = 0
		for i := 0; i < 1000; i++ {
			go inc(&x)
		}
		fmt.Println("x[", h, "]:", x)
	}
}
実行結果
x[ 0 ]: 981
x[ 1 ]: 966
x[ 2 ]: 910
x[ 3 ]: 943
x[ 4 ]: 1000
```
実行結果をみてわかるように、5回全て結果が異なる。本来は全て1000になるべきだ。この違いの原因は、インクリメントのタイミングである。このプログラムでは1000回のインクリメントをgoroutineで同時に実行している。簡単にインクリメントと言っているが、実際は
- 値が格納されているメモリを参照する
- 現在の値を取得する
- 現在の値に１加える
- １加えた値を元のメモリに格納する
という作業を行っている。環境によってそれぞれのインクリメントのタイミングが多少前後するが1000回が一斉に実行されるので中には値を取得するタイミング被ったりし、インクリメントが反映されない状況が出てくる。
<br>
そのような自体を防ぐために、goroutineどうしでコミュニケーションを取り合うシステムがChannelである。

```
func inc(x *int, ch chan string) {
	*x++
	ch <- "done"
}

func main() {
	var x int
	ch := make(chan string)
	for h := 0; h < 5; h++ {
		x = 0
		for i := 0; i < 1000; i++ {
			go inc(&x, ch)
			<-ch
		}
		fmt.Println("x[", h, "]:", x)
	}
}
実行結果
x[ 0 ]: 1000
x[ 1 ]: 1000
x[ 2 ]: 1000
x[ 3 ]: 1000
x[ 4 ]: 1000
```

<br>
こう言ったものをメッセージパッシングとかいうらしい。詳しくはメッセージパッシングとかCSP(Communicasing Sequential Processes)とかでググってもらいたい。

channelの特徴
- メッセージの型を指定できる
- first class value であり、引数や戻り値にも使える
- send/receive でブロックする
- buffer で、一度に扱えるメッセージ量を指定できる

channelはmake(chan 型)で宣言し <- で受信・送信を行う。<br>
ch <- :Sender チャンネルに送信<br>
<- ch :Reciever チャンネルから受信<br>
以下のコードと実行結果を見れば、(1)でch自身を出力しているのを見れば、チャンネルはポインターのようにメモリのアドレスを指している。
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
複数の値を送信から受け取っている場合はrangeでループを回すことが出来る。この時ループはチャンネルが閉じられているところまで繰り返される。なのでrangeで回す時は、その前にclose()で必ず閉じておく必要がある。閉じていなかった場合は、パニックを起こします。
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
注：チャネルを閉じる必要があるのは送信側だけで、受信側は閉じないでください。閉じたチャネルに送信すると、パニックが発生します。
<br>
別のメモ：チャネルはファイルとは異なります。通常は閉じる必要はありません。閉じる必要があるのは、rangeループを終了するなど、受信側に値が来ないことを通知する必要がある場合のみです。

## Select
select文はswitch文と似ていますがswitch文と異なり、チャンネルが送信できるか受信できるかによって操作を分岐する文です。goroutineが複数の通信操作を待機できるようになります。
selectは、そのケースの1つが実行可能になるまでブロックし、その後そのケースを実行します。複数の準備ができている場合は、ランダムに1つを選択します。
<br />
<br />
下の例では以前やったフィボナッチ数列をgoroutine, channel, selectを使ってやっている。コードからは処理の流れが読み取りにくいが次のようなことをやっている。
<br />
main()

1. チャンネルc, quitの定義
2. 即時関数の実行
3. fibonacci()の実行:quitは何も受信してないため、quitから送信できない。cは受信可能なので*2が実行される
4. cからの送信が可能になるため即時関数内のループのi=0回目を実行し、Println()でcから受信した値を表示
5. 3.を実行
6. 4.を実行
以後これをi=9まで合計10回繰り返す。
7. 即時関数のforループが終了したので、quitが0を受信する。
8. fibonacci()の実行:quitが0を受信しているので*1が実行され、returnでfibonacci()は終了される。

<br />

```
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case <-quit: // *1
			time.Sleep(200 * time.Millisecond)
			fmt.Println("quit")
			return
		case c <- x: // *2
			time.Sleep(200 * time.Millisecond)
			x, y = y, x+y
		}
	}
}

func main() {
	fmt.Println("go start.")
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
/*
実行結果
$ ./select 
go start.
0
1
略
34
quit
*/
```
以下の場合だと、fibonacci()を実行後に即時関数を実行する手順になる。fibonacchi()でチャンネルcから受信する必要があるが、cの中身はからなのでdeadlockが発生しエラーになる。並行処理の後に持ってくる必要がある。
```
func main() {
	c := make(chan int)
	quit := make(chan int)

	fmt.Println("go start.")

	fibonacci(c, quit)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

}
/*
実行結果
$ ./select 
go start.
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [select]:
main.fibonacci(0xc0000200c0, 0xc000020120)
        /go/src/tour_of_go/channel/select/select.go:11 +0xe8
main.main()
        /go/src/tour_of_go/channel/select/select.go:29 +0xd7
*/
```
以下のようにfibonacci関数側ではなく即時関数側にselect文を書いて同じ実装も可能。
```
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for i := 0; i < 10; i++ {
		time.Sleep(200 * time.Millisecond)
		c <- x
		x, y = y, x+y
	}
	quit <- 0
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	fmt.Println("go start.")

	go fibonacci(c, quit)
	func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-quit:
				fmt.Println("quit")
				return
			}
		}
	}()
}
/*
実行結果(略)
*/
```

defaultを使うと、他のケースの準備ができていないときに行う処理を実装できる。以下の例では1秒ごとにtick.を表示し3秒後にBOOM!を表示する。それ以外の場合は0.5秒感覚で"    ."を表示する(defaultの処理)。<br/>
このようにチャンネルから受信・送信するまでに別の処理をさせておくこともできる。
```
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1000 * time.Millisecond)
	boom := time.After(3000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
/*
実行結果
    .
    .
tick.
    .
    .
tick.
    .
    .
BOOM!
(3秒をカウント！)
*/
```
引用:[a tour of go](https://tour.golang.org/concurrency/6)

## sync.Mutex
並行処理でのゴルーチン間のコミュニケーションにチャンネルが有効だというのはこれまでの内容でわかった。
しかし、コミュニケーションを必要としないが処理の衝突を避けたい場合がある。その時使うのがmutexである。
mutexによってロックされている間は他の処理が実行されない。なので以下の例は複数のゴルーチンが同時にインクリメントを実行し値がその時々で変化することはない。ただし、これはこれはロックを行っているinc関数内での話であり、fmt.Printlnは関係ない。なのでPrintlnで表示する前にtime.Sleepで遅延をおこす必要がある。
```
func inc(x *int, m *sync.Mutex) {
	m.Lock()
	*x++
	m.Unlock()
}

func main() {
	var x int
	var m sync.Mutex
	for h := 0; h < 5; h++ {
		x = 0
		for i := 0; i < 1000; i++ {
			go inc(&x, &m)
		}
		time.Sleep(time.Millisecond)
		fmt.Println("x[", h, "]:", x)
	}
}
実行結果
x[ 0 ]: 1000
x[ 1 ]: 1000
x[ 2 ]: 1000
x[ 3 ]: 1000
x[ 4 ]: 1000
```
time.Sleepで遅延を起こさずやるならば、sync.WaitGroupで処理が終わるまでPrintlnの実行を待たせる方法がある。sync.WaitGroupの書き方は範囲外なのでここでは触れない。

```
func inc(x *int, m *sync.Mutex, w *sync.WaitGroup) {
	m.Lock()
	*x++
	m.Unlock()
	w.Done()
}

func main() {
	var x int
	var m sync.Mutex
	var w sync.WaitGroup
	for h := 0; h < 5; h++ {
		x = 0
		for i := 0; i < 1000; i++ {
			w.Add(1)
			go inc(&x, &m, &w)
		}
		w.Wait()
		fmt.Println("x[", h, "]:", x)
	}
}
実行結果
x[ 0 ]: 1000
x[ 1 ]: 1000
x[ 2 ]: 1000
x[ 3 ]: 1000
x[ 4 ]: 1000
```

