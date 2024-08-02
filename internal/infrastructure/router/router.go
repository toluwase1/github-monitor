package router

import (
	"github-monitor/internal/interface/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(ctrl *controller.Controller) *gin.Engine {
	r := gin.Default()

	r.GET("/repositories/:name/commits", ctrl.GetCommits)
	r.GET("/repositories/:name", ctrl.GetRepository)
	r.POST("/repositories/:name/reset", ctrl.ResetCollection)
	r.GET("/authors/top/:n", ctrl.GetTopAuthors)
	return r
}

//test
//config date
//separation of concern
