package main

import (
	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.Run()
}

/*
1- Write in Terminal:
	go mod tidy

2- Output:
	go: finding module for package github.com/beego/beego/v2/server/web
	go: downloading github.com/beego/beego/v2 v2.1.3
	go: found github.com/beego/beego/v2/server/web in github.com/beego/beego/v2 v2.1.3
	go: downloading github.com/prometheus/client_golang v1.16.0
	go: downloading golang.org/x/crypto v0.10.0
	go: downloading google.golang.org/protobuf v1.30.0
	go: downloading github.com/elazarl/go-bindata-assetfs v1.0.1
	go: downloading github.com/stretchr/testify v1.8.1
	go: downloading github.com/pkg/errors v0.9.1
	go: downloading github.com/mitchellh/mapstructure v1.5.0
	go: downloading github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18
	go: downloading github.com/prometheus/client_model v0.3.0
	go: downloading github.com/prometheus/common v0.42.0
	go: downloading github.com/cespare/xxhash/v2 v2.2.0
	go: downloading github.com/prometheus/procfs v0.10.1
	go: downloading golang.org/x/net v0.10.0
	go: downloading github.com/golang/protobuf v1.5.3
	go: downloading github.com/matttproud/golang_protobuf_extensions v1.0.4
	go: downloading golang.org/x/text v0.10.0
	go: downloading gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c
	go: downloading github.com/kr/pretty v0.3.1
	go: downloading github.com/rogpeppe/go-internal v1.10.0
	go: finding module for package github.com/kr/text
	go: downloading github.com/kr/text v0.2.0
	go: found github.com/kr/text in github.com/kr/text v0.2.0

*/

/*
1- After that, write:
	go run .

2- Output in Terminal is:
	2023/11/11 08:10:26.835 [D]  init global config instance failed. If you do not use this, just ignore it.  open conf/app.conf: no such file or directory
	2023/11/11 08:10:26.839 [I]  http server Running on http://:8080

3- write in browser:
	localhost:8080

4- Output:
	NOT Found blue text with another information.
*/
