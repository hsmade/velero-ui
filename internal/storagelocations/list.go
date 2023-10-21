package storagelocations

import (
	"context"
	"github.com/hsmade/velero-ui/internal/k8s"
	velerov1api "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

// ListStorageLocations returns all found `storagelocations` custom resources
func ListStorageLocations(ctx context.Context, dynamicClient dynamic.Interface) ([]velerov1api.BackupStorageLocation, error) {
	var storagelocations []velerov1api.BackupStorageLocation
	err := k8s.ListByFilter(
		ctx,
		dynamicClient,
		schema.GroupVersionResource{
			Group:    "velero.io",
			Version:  "v1",
			Resource: "backupstoragelocations",
		},
		metav1.ListOptions{},
		&storagelocations,
	)
	return storagelocations, err
}
