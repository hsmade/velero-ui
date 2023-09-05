package schedules

import (
	"context"
	"fmt"
	"github.com/hsmade/velero-ui/util"
	velerov1api "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

type ScheduleDetail struct {
	Schedule *velerov1api.Schedule
	Backups  []*velerov1api.Backup
}

func Detail(ctx context.Context, dynamicClient dynamic.Interface, name string) (*ScheduleDetail, error) {
	schedule, err := getSchedule(ctx, dynamicClient, name)
	if err != nil {
		return nil, fmt.Errorf("get schedule: %w", err)
	}

	backups, err := getBackups(ctx, dynamicClient, schedule.Name)
	if err != nil {
		return nil, fmt.Errorf("get podvolume backups: %w", err)
	}

	result := ScheduleDetail{
		Schedule: schedule,
		Backups:  backups,
	}

	return &result, nil
}

func getSchedule(ctx context.Context, dynamicClient dynamic.Interface, name string) (*velerov1api.Schedule, error) {
	schedule := new(velerov1api.Schedule)
	err := util.GetByName(ctx, dynamicClient, schema.GroupVersionResource{
		Group:    "velero.io",
		Version:  "v1",
		Resource: "schedules",
	},
		name,
		&schedule,
	)
	return schedule, err
}

func getBackups(ctx context.Context, dynamicClient dynamic.Interface, name string) ([]*velerov1api.Backup, error) {
	backups := new([]*velerov1api.Backup)
	err := util.ListByFilter(ctx, dynamicClient,
		schema.GroupVersionResource{
			Group:    "velero.io",
			Version:  "v1",
			Resource: "backups",
		},
		metav1.ListOptions{LabelSelector: fmt.Sprintf("velero.io/schedule-name=%s", name)},
		backups,
	)
	return *backups, err
}
