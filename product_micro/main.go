package main

import (
	"Ecommerce3/approduct"
)

func main() {
	App := &approduct.App{}
	App.Initialize()
	App.Run(":5000")

}