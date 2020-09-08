# Images
このエクササイズでは、画像ストリームを扱います。配列のところで画像を出力しましたが、あの時はa tour of go側が用意していたメソッドを利用して配列をImageに変換し出力していました。今回はそのImageオブジェクトを扱うImageインターフェースを元にImageを生成しようというエクササイズです。Imageインターフェースを扱うのでStringer, Error, Readerの続きとも言えます。
```
package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```

今回はImageというオブジェクトを定義し、そこに上であげているようなImageの生成に必要となる３つのメソッドを定義します。ColorModel()はcolor.RGBAModel, Bounds()はimage.Rect, At(x, y int)はcolor.RGBAで実装します。

### ColorModel()
カラーモデルを返します。例えばRGBAなど。
他にも Alpha, CMYK, YCbCr(輝度, 青系統, 赤系統)がある。
### Bounds()
At()が色を返すことのできる領域を返します。言い換えるとBounds()で図形の形を決めることが出来ます。
### At(x, y int)
各ピクセル(x, y)の色をを返す。

それぞれたくさんのメソッドが用意されているので詳しくは[公式のドキュメント](https://golang.org/pkg/image/)を参照してください。
分からない人はRBA関係のだけでもみておくと理解が深まると思います。

文字列を表示するコンソールに出力するためには"golang.org/x/tour/pic"のpic.ShowImage(m)を使用します。
以下のA tour of goのコンソールで試してください。
[Exercise: Images](https://tour.golang.org/methods/25)