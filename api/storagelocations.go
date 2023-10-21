package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hsmade/velero-ui/internal/storagelocations"
	"log/slog"
	"net/http"
)

func (a *Api) ListStorageLocations(c *gin.Context) {
	result, err := storagelocations.ListStorageLocations(c, a.dynamicClient)
	if err != nil {
		slog.Error("listing storagelocations", "err", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func (a *Api) GetStorageLocation(c *gin.Context) {
	name := c.Param("name")
	result, err := storagelocations.GetStorageLocationDetail(c, a.dynamicClient, name)
	if err != nil {
		slog.Error("getting storagelocation", "err", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}
