package router

import (
	"github-monitor/internal/interface/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(ctrl *controller.Controller) *gin.Engine {
	r := gin.Default()

	r.GET("/repositories/:name/commits", ctrl.GetCommits) //1
	r.GET("/repositories/:name", ctrl.GetRepository)
	r.GET("/reset", ctrl.ResetCollection) //2
	//author with greates number of commits
	return r
}
