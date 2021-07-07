package main

import (
	"Ecommerce4/appayment"
)

func main() {
	App := &appayment.App{}
	App.Initialize()
	App.Run(":6000")

}