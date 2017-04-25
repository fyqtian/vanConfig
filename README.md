自学GO写的配置加载目前只支持JSON

需要github.com/bitly/go-simplejson

安装
go get github.com/bitly/go-simplejson

go get github.com/fyqtian/vanConfig

##Quick Start

######Create file `hello.go`
```go
package main

import "github.com/fyqtian/vanConfig"

func main(){
   configJson := vanConfig.Config{FilePath:"E:/go/2.json",FileType:"json"}

   	err := configJson.Open()

   	if err != nil{
   		fmt.Print(err)
   		return
   	}

   	id := configJson.ParseJson.Get("require").Get("laravel/framework")

   	fmt.Println(id)
}
```