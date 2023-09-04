package backups

import (
	"context"
	"github.com/hsmade/velero-ui/util"
	velerov1api "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

func List(ctx context.Context, dynamicClient dynamic.Interface) ([]velerov1api.Backup, error) {
	var backups []velerov1api.Backup
	err := util.ListByFilter(
		ctx,
		dynamicClient,
		schema.GroupVersionResource{
			Group:    "velero.io",
			Version:  "v1",
			Resource: "backups",
		},
		metav1.ListOptions{},
		&backups,
	)
	return backups, err
}
