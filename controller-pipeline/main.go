package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/pipeline/:type", func(c *gin.Context) {
		pipelineType := c.Param("type")
		if pipelineType == "simple" {
			normalBuilder := getBuilder(pipelineType)
			director := newDirector(normalBuilder)
			res := director.buildResponse()
			c.JSON(http.StatusOK, res)
		} else if pipelineType == "simpleAsr" {
			simpleAsrBuilder := getSimpleAsrBuilder(pipelineType)
			director := newAsrDirector(simpleAsrBuilder)
			res := director.buildAsrResponse()
			c.JSON(http.StatusOK, res)
		}
	})
	r.Run()

}
