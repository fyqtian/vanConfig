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
   config := &vanConfig.Config{FilePath: "2.txt", FileType: "text"}
   	err := vanConfig.NewConfig(config)
   	if err != nil {
   		fmt.Println(err)
   	}
   	fmt.Println(config.ParseJson.Get("id").Get("van").String())
}
```

```json
{
  "name": "webclient/runchina",
  "description": "The set of pages used at pptv.",
  "keywords": ["runchina.pptv.com"],
  "license": "MIT",
  "require": {
    "laravel/framework": "4.2.*",
    "guzzlehttp/guzzle": "5.*",
    "webclient/synacast": "1.3.5",
    "phpspec/phpspec": "~2.1",
    "maatwebsite/excel": "1.*"
  },
  "minimum-stability": "stable"
}
```

```text
[id]
name=van
env=test

[where]
home=gucun
```