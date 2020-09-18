# Web Crawler
これで最後のエクササイズです。最後はウェブクローラーを修正して並行処理出来るようにしてください。という問題です。
すでにプログラムのコードが与えたれていると思いますが、これだけでもきちんと動作します。しかし、urlが重複しているものもありますし、何より逐次処理になっています。なのでurlが重複しているものは除外して、かつ並行処理が出来るようにCrawl関数を書き換えて欲しいというわけです。

少しコードの説明をしておきます。

Fetch インターフェースです。後述するFetcherのためのものです。

```
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}
```

書き換えて欲しいCrawl関数です。必要があれば二つ以上の関数に分けても構いません。
```
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}
```

main関数
```
func main() {
	Crawl("https://golang.org/", 4, fetcher)
}
```

Fetcher。本来ならクローリングはWeb上からデータを取得してくるべきでしょうがセキュリティなどの問題から許可なくクローリングすることは許されていません。なのでここでクローリングで取得するデータを用意しています。
ここの内容が重複なく取得できればOKです。ちなみに重複なく取得するとデータは５つになります。
```
// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
```

今回のエクササイズは並行処理なのでゴルーチンもちろん、これまでの総まとめなのでこれまでに使ったものを駆使してください。それでも難しいと思います。その時はsync.WaitGroupを調べて使ってみるのも手です。
私が残している答えはsync.WaitGroupを使っています。
[Exercise: Web Crawler](https://tour.golang.org/concurrency/10)