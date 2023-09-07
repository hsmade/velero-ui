package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hsmade/velero-ui/internal/backups"
	"log/slog"
	"net/http"
)

func (a *Api) ListBackups(c *gin.Context) {
	result, err := backups.ListBackups(c, a.dynamicClient)
	if err != nil {
		slog.Error("listing backups", "err", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func (a *Api) GetBackup(c *gin.Context) {
	name := c.Param("name")
	result, err := backups.GetBackupDetail(c, a.dynamicClient, name)
	if err != nil {
		slog.Error("getting backup", "err", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}
