package main

import (
	sliceoapi "github.com/cglLaLaLa/SoVITS_web/common/client/slice"
	"github.com/cglLaLaLa/SoVITS_web/preprocess/app"
	"github.com/cglLaLaLa/SoVITS_web/preprocess/app/command"
	"github.com/cglLaLaLa/SoVITS_web/preprocess/convertor"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	app app.Application
}

func (s HttpServer) PostPreprocessCustomerIdSlice(c *gin.Context, customerId string) {
	var (
		req  sliceoapi.SliceRequest
		err  error
		resp struct {
			CustomerID  string `json:"customer_id"`
			RedirectURL string `json:"redirect_url"`
		}
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		return
	}
	r, err := s.app.Commands.StartSliceCommand.Handle(c.Request.Context(), command.StartSlice{
		CustomerID: customerId,
		Request:    convertor.NewSliceStartConverter().ClientToEntity(&req),
	})
	if err != nil {
		return
	}
	resp.CustomerID = customerId
	//todo
	resp.RedirectURL = "todo" + r.SliceId
}
