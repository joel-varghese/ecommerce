package main 

import (
	"appuser"
)

func main() {
	App := &appuser.App{}
	App.Initialize()
	App.Run(":3000")

}