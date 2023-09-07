package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hsmade/velero-ui/internal/schedules"
	"log/slog"
	"net/http"
)

func (a *Api) ListSchedules(c *gin.Context) {
	result, err := schedules.ListSchedules(c, a.dynamicClient)
	if err != nil {
		slog.Error("listing schedules", "err", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func (a *Api) GetSchedule(c *gin.Context) {
	name := c.Param("name")
	result, err := schedules.GetScheduleDetail(c, a.dynamicClient, name)
	if err != nil {
		slog.Error("getting schedule", "err", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}
