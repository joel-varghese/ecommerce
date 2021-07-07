package main

import (
	"appcart"
)

func main() {
	App := &appcart.App{}
	App.Initialize()
	App.Run(":4000")

}