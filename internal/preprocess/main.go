package main

import (
	"context"
	"github.com/cglLaLaLa/SoVITS_web/common/server"
	sliceoapi "github.com/cglLaLaLa/SoVITS_web/preprocess/ports"
	"github.com/cglLaLaLa/SoVITS_web/preprocess/service"
	"github.com/gin-gonic/gin"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	application, cleanup := service.NewApplication(ctx)
	defer cleanup()
	server.RunHttpServer("sovits", func(router *gin.Engine) {
		//router.StaticFile("/success", "../../public/success.html")
		sliceoapi.RegisterHandlersWithOptions(router, HttpServer{
			app: application,
		}, sliceoapi.GinServerOptions{
			BaseURL:      "/api",
			Middlewares:  nil,
			ErrorHandler: nil,
		})
	})
}
