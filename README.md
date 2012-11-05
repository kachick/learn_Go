My notebooks of Go
==================

Install
--------

http://golang.jp/install

Windows�ł́Azip���𓀂��Ăǂ����ւ����A���̓�ݒ�
* GOROOT���ϐ��� \go
* PATH���ϐ��� \go\bin


�R���p�C���Ƃ�
--------------

* ��ɓ��

```bash
go run foo.go
```

```bash
go tool 8g foo.go
go tool 8l foo.8
foo.out.exe
```

* i386����8g, 8l���g�����ɂȂ�Aamd64�Ȃ�6g, 6l���g�����ɂȂ�B


�\���Ƃ�
-------

����Hello�ŁA���\�ȃ|�C���g������

```go
package main

import "fmt"

func main() {
    fmt.Printf("Hello\n")
}
```

* package �ɂ�""�Ƃ����Ȃ�
* import �Ń��C�u�����ǂݍ��ނ��A�����ł�""����
* func�Ŋ֐���`
* �s���͉��s�݂̂�OK�B�R���p�C��������ɉ��s��ǉ�����
* ���C�u������.foo �ŁA���̃��C�u�������̊e�v�f�փA�N�Z�X����B

���C�u�������̋�����������i�Ƃ��āA���܂萄������Ȃ�����import���ɖ��O��ς����i������

```go
package main

import format "fmt"

func main() {
    format.Printf("Hello\n")
}
```

�^�錾�ȗ����Ȃ���ϐ��ւ����ނƁA����

```go
package main

import "fmt"

func main() {
    str := "hello, world\n"
    fmt.Printf(str)
}
```

�Q�l
----

### Tips

* [����d�l](http://golang.jp/go_spec)
* [Effective Go](http://go.shibu.jp/effective_go.html)
* [�v���O���~���O����Go�t���[�Y�u�b�N](http://www.amazon.co.jp/exec/obidos/ASIN/486401096X/yoshikisbooks-22/ref=nosim)
* [���A����](http://www.informit.com/store/product.aspx?isbn=0321817141)
* [���A�T���v���R�[�h](http://www.informit.com/content/images/9780321817143/downloads/9780321817143_codeexamples.zip)

### Tutorial

* [A Tour of Go](http://go-tour-jp.appspot.com/#1)
* [go-tour](http://code.google.com/p/go-tour/)

### Others

* [golang.jp](http://golang.jp)

Author
------

Kenichi Kamiya
