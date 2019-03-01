What is wasm-loader?

https://github.com/vuejs/vue-cli/issues/763
* https://github.com/greenpdx/vueopencvjs/blob/master/src/components/HelloWorld.vue
* https://github.com/greenpdx/vueopencvjs/blob/master/src/components/add.es6

How to generate a wasm file with go

1. write a xx.fo file 
2. then run command
`GOARCH=wasm GOOS=js go build -o calc.wasm calc.go`

if there is no wasm_exec.html,js in rootPath 
you can use this 
`cp $(go env GOROOT)/misc/wasm/wasm_exec.{html,js} .`
copy those from go env

build a web-server

http.go
```
package main
import (
    "flag"
    "log"
    "net/http"
    "strings"
)

var (
    listen = flag.String("listen", ":8080", "listen address")
    dir    = flag.String("dir", ".", "directory to serve")
)

func main() {
    flag.Parse()
    log.Printf("listening on %q...", *listen)
    log.Fatal(http.ListenAndServe(*listen, http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
        if strings.HasSuffix(req.URL.Path, ".wasm") {
            resp.Header().Set("content-type", "application/wasm")
        }

        http.FileServer(http.Dir(*dir)).ServeHTTP(resp, req)
    })))
}
```

`go run http.go ` 

you will see the results
