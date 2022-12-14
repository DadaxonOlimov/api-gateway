package api

import (
	"github.com/DadaxonOlimov/services"
	proto2 "github.com/golang/protobuf/proto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
)

type Controller struct {
	Services services.ServiceManager
}

func NewController(services services.ServiceManager) *Controller {
	return &Controller{
		Services: services,
	}
}

type Message struct {
	Message string `json:"message" example:"message"`
}

func (h *Controller) ResponseProtoJson(c *gin.Context, p proto2.Message) {
	var (
		jsonbMarshal jsonpb.Marshaler
	)

	jsonbMarshal.OrigName = true
	jsonbMarshal.EmitDefaults = true

	js, err := jsonbMarshal.MarshalToString(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}