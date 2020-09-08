# rot13Readers
ここでは引き続きio.Readerインターフェースを使用してRead()関数を書き換えて与えられた暗号文を復号化してみましょう。有名な暗号化の手法にはシーザー暗号があります。

シーザー暗号はa -> d, e -> h, z -> cのようにアルファベット順に３文字ずらして文章を暗号化します。
例えば、

thank you -> wkdqn brx

のようになります。今回のrot13もシーザー暗号の一種で13文字ずらして暗号文を生成します。

### strings.NewReader(s string) *Reader
NewReaderは、sから読み込みを行う新しいReaderを返します。

### io.Copy func Copy(dst Writer, src Reader) (written int64, err error)
srcからdstにコピーを行います。EOFかエラーが起きるまでコピーします。コピーしたバイト数とエラーが起きた場合エラーを返します。

### os.Stdout
標準出力を行う関数。

ここからはまだ未解決で予想です。
- main()の処理について
    - string.NewReader()は文字列からio.Readerに渡す時に使う関数。これでrot13Readerのプロパティとして使える。
    - これを変数rに代入する。これで内部で持っている文字列はRead()関数によって出力する際はrot13でアルファベット順に13文字シフトした文字が出力される。
    - io.Copyはオブジェクトを標準出力os.Stdoutにコピーすることで表示コンソールに表示している。
    - fmt.Printを使っていないのはbyte列で操作をしているからだと予想。

[Exercise: rot13Reader](https://tour.golang.org/methods/23)