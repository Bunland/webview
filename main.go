package main

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"
import "github.com/webview/webview"

//export CreateWindow
func CreateWindow() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Basic Example")
	w.SetSize(480, 320, webview.HintNone)
	w.SetHtml("Thanks for using webview!")
	w.Run()
}

func main() {}
