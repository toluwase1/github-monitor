package controller

import (
	"github-monitor/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	commitUsecase     usecase.CommitUsecase
	repositoryUsecase usecase.RepositoryUsecase
}

func NewController(commitUC usecase.CommitUsecase, repoUC usecase.RepositoryUsecase) *Controller {
	return &Controller{
		commitUsecase:     commitUC,
		repositoryUsecase: repoUC,
	}
}

func (ctrl *Controller) GetCommits(c *gin.Context) {
	repoName := c.Param("name")
	commits, err := ctrl.commitUsecase.GetCommitsByRepositoryName(repoName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, commits)
}

func (ctrl *Controller) GetRepository(c *gin.Context) {
	repoName := c.Param("name")
	repo, err := ctrl.repositoryUsecase.GetRepositoryByName(repoName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, repo)
}
