package controller

import (
	"github-monitor/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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
	commits, err := ctrl.commitUsecase.GetCommitsByRepoNameFromDB(repoName)
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

func (ctrl *Controller) ResetCollection(c *gin.Context) {
	repoName := c.Param("name")
	startDateStr := c.Query("start_date")
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format. Use YYYY-MM-DD."})
		return
	}

	err = ctrl.commitUsecase.ResetCollection(repoName, startDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Collection reset successfully."})
}

func (ctrl *Controller) GetTopAuthors(c *gin.Context) {
	n, err := strconv.Atoi(c.Param("n"))
	if err != nil || n <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number of authors specified."})
		return
	}

	topAuthors, err := ctrl.commitUsecase.GetTopAuthorsByCommitCount(n)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, topAuthors)
}
