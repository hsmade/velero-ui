package backups

import (
	"context"
	"fmt"
	velerov1api "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

func List(ctx context.Context, dynamicClient dynamic.Interface) ([]*velerov1api.Backup, error) {
	crList, err := dynamicClient.
		Resource(schema.GroupVersionResource{
			Group:    "velero.io",
			Version:  "v1",
			Resource: "backups",
		}).
		Namespace("velero").
		List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("listing backups: %w", err)
	}

	var backups []*velerov1api.Backup

	for _, item := range crList.Items {
		backup := new(velerov1api.Backup)
		err = runtime.DefaultUnstructuredConverter.FromUnstructured(item.UnstructuredContent(), &backup)
		if err != nil {
			return nil, fmt.Errorf("unmarshal backup resource: %w", err)
		}
		backups = append(backups, backup)
	}
	return backups, nil
}
