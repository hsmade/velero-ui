package schedules

import (
	"context"
	"github.com/hsmade/velero-ui/util"
	velerov1api "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

func List(ctx context.Context, dynamicClient dynamic.Interface) ([]velerov1api.Schedule, error) {
	var schedule []velerov1api.Schedule
	err := util.ListByFilter(
		ctx,
		dynamicClient,
		schema.GroupVersionResource{
			Group:    "velero.io",
			Version:  "v1",
			Resource: "schedules",
		},
		metav1.ListOptions{},
		&schedule,
	)
	return schedule, err
}
