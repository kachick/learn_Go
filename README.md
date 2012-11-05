My notebooks of Go
==================

Install
--------

http://golang.jp/install

Windows版は、zipを解凍してどこかへおき、次の二つ設定
* GOROOT環境変数へ \go
* PATH環境変数へ \go\bin


コンパイルとか
--------------

* 主に二つ

```bash
go run foo.go
```

```bash
go tool 8g foo.go
go tool 8l foo.8
foo.out.exe
```

* i386だと8g, 8lを使う事になり、amd64なら6g, 6lを使う事になる。


構文とか
-------

次のHelloで、結構なポイントがある

```go
package main

import "fmt"

func main() {
    fmt.Printf("Hello\n")
}
```

* package には""とかつけない
* import でライブラリ読み込むが、ここでは""つける
* funcで関数定義
* 行末は改行のみでOK。コンパイラが勝手に改行を追加する
* ライブラリ名.foo で、そのライブラリ内の各要素へアクセスする。

ライブラリ名の競合を避ける手段として、あまり推奨されない物のimport時に名前を変える手段もある

```go
package main

import format "fmt"

func main() {
    format.Printf("Hello\n")
}
```

型宣言省略しながら変数へつっこむと、こう

```go
package main

import "fmt"

func main() {
    str := "hello, world\n"
    fmt.Printf(str)
}
```

参考
----

### Tips

* [言語仕様](http://golang.jp/go_spec)
* [Effective Go](http://go.shibu.jp/effective_go.html)
* [プログラミング言語Goフレーズブック](http://www.amazon.co.jp/exec/obidos/ASIN/486401096X/yoshikisbooks-22/ref=nosim)
* [同、原著](http://www.informit.com/store/product.aspx?isbn=0321817141)
* [同、サンプルコード](http://www.informit.com/content/images/9780321817143/downloads/9780321817143_codeexamples.zip)

### Tutorial

* [A Tour of Go](http://go-tour-jp.appspot.com/#1)
* [go-tour](http://code.google.com/p/go-tour/)

### Others

* [golang.jp](http://golang.jp)

Author
------

Kenichi Kamiya
