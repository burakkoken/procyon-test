package controller

import web "github.com/procyon-projects/procyon-web"

type HelloRequest struct {
	PathVariables struct {
		Name string `json:"name" yaml:"name"`
	} `request:"path"`
}

type ImpHelloController struct {
}

func NewHelloController() ImpHelloController {
	return ImpHelloController{}
}

func (controller ImpHelloController) RegisterHandlers(registry web.HandlerRegistry) {
	registry.Register(
		web.Get(controller.Hello, web.Path("/hello/:name"), web.RequestObject(HelloRequest{})),
	)
}

func (controller ImpHelloController) Hello(ctx *web.WebRequestContext) {
	helloRequest := &HelloRequest{}
	if err := ctx.BindRequest(helloRequest); err != nil {
		ctx.ThrowError(web.HttpErrorBadRequest)
	}
	ctx.SetModel("Hello " + helloRequest.PathVariables.Name)
	ctx.SetResponseContentType(web.MediaTypeApplicationTextHtml)
}
