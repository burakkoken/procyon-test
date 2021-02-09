package main

import (
	"github.com/burakkoken/procyon-test/controller"
	core "github.com/procyon-projects/procyon-core"
)

func init() {
	core.Register(controller.NewHelloController)
}
