package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

type Api struct {
	dynamicClient dynamic.Interface
}

func New(router *gin.RouterGroup) error {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	config, err := kubeConfig.ClientConfig()
	if err != nil {
		return fmt.Errorf("setting up kubernetes config: %w", err)
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("setting up kubernetes dynamic client: %w", err)
	}

	api := Api{dynamicClient: dynamicClient}
	router.GET("/schedules", api.ListSchedules)
	router.GET("/schedule/:name", api.GetSchedule)
	router.GET("/backups", api.ListBackups)
	router.GET("/backup/:name", api.GetBackup)
	return nil
}
