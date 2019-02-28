package main

import (
    "flag"
    "log"
    "net/http"
)

func main() {
    port := flag.String("port", "8000", "port")
    flag.Parse()
    http.HandleFunc("/", rootHandle)
    http.HandleFunc("wasm",wasmHandle)
    log.Fatal(http.ListenAndServe(":"+*port, http.FileServer(http.Dir("."))))
}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, page, html.EscapeString(hex.Dump(wasmAdd)))
}