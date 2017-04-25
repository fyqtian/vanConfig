自学GO写的配置加载目前只支持JSON

需要github.com/bitly/go-simplejson

安装
go get github.com/bitly/go-simplejson

##Quick Start

######Create file `hello.go`
```go
package main

import "github.com/astaxie/beego"

func main(){
    beego.Run()
}
```