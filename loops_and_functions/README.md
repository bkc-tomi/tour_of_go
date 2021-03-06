# Loops and Functions
ここではif文とfor文を使用して平方根を求める関数を定義します。
定義した関数を使用して平方根を求めます。
平方根を求める際には[ニュートン法](https://ja.wikipedia.org/wiki/%E3%83%8B%E3%83%A5%E3%83%BC%E3%83%88%E3%83%B3%E6%B3%95)を使用します。
これは「<img src="https://latex.codecogs.com/gif.latex?y=x^2" />のような関数の折線とx軸との交点を求めそこでの折線を求める。」といった作業を繰り返すと、関数とx軸との交点に近づいていく性質を利用して平方根を求めます(詳しくはニュートン法のリンクを)。
理屈はさておき、折線とx軸との交点は以下の式で求まります。

z -= (z*z - x) / (2*z)

これを繰り返すと平方根が求まるので、for文が必要になります。
また平方根は正の値しか存在しないので、if文で正の値のみを処理するようにしてあげる必要があります。

[Exercise: Loops and Functions](https://tour.golang.org/flowcontrol/8)