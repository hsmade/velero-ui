package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hsmade/velero-ui/internal/schedules"
	"net/http"
)

func (a *Api) ListSchedules(c *gin.Context) {
	result, err := schedules.List(c, a.dynamicClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func (a *Api) GetSchedule(c *gin.Context) {
	name := c.Param("name")
	result, err := schedules.Detail(c, a.dynamicClient, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}
