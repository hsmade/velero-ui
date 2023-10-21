package storagelocations

import (
	"context"
	"fmt"
	"github.com/hsmade/velero-ui/internal/k8s"
	velerov1api "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

type StorageLocationDetail struct {
	StorageLocation *velerov1api.BackupStorageLocation `json:"storage_location"`
}

// GetStorageLocationDetail returns a struct with the backup info for the backup specified by `name`
func GetStorageLocationDetail(ctx context.Context, dynamicClient dynamic.Interface, name string) (*StorageLocationDetail, error) {
	storagelocation, err := getStorageLocation(ctx, dynamicClient, name)
	if err != nil {
		return nil, fmt.Errorf("get storagelocation: %w", err)
	}

	result := StorageLocationDetail{
		StorageLocation: storagelocation,
	}

	return &result, nil
}

// getStorageLocation gets a single storagelocation by `name`
func getStorageLocation(ctx context.Context, dynamicClient dynamic.Interface, name string) (*velerov1api.BackupStorageLocation, error) {
	storagelocation := new(velerov1api.BackupStorageLocation)
	err := k8s.GetByName(ctx, dynamicClient, schema.GroupVersionResource{
		Group:    "velero.io",
		Version:  "v1",
		Resource: "backupstoragelocations",
	},
		name,
		&storagelocation,
	)
	return storagelocation, err
}
