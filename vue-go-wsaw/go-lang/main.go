package main

import (
	"syscall/js"
)

func changeBodyColor(color []js.Value){
	js.Global().Get("document").Get("body").Set("style","color:"+color[0].String())
	print("have", color[0].String())
}

func setInputValue(val []js.Value) {
	print("have", val)
	id := val[0].String()
	// document.getElementById(id).value = "value from Go"
	js.Global().Get("document").Call("getElementById", id).Set("value", "value from Go")
}

// 将Go里面的方法注入到window.fibNative里面
func registerCallbacks() {
	js.Global().Set("changeBodyColor", js.NewCallback(changeBodyColor))
	js.Global().Set("setInputValue", js.NewCallback(setInputValue))
}

func main() {
	registerCallbacks()
	select {}
}