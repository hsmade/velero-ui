package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hsmade/velero-ui/internal/backups"
	"net/http"
)

func (a *Api) ListBackups(c *gin.Context) {
	result, err := backups.List(c, a.dynamicClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func (a *Api) GetBackup(c *gin.Context) {
	name := c.Param("name")
	result, err := backups.Detail(c, a.dynamicClient, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}
