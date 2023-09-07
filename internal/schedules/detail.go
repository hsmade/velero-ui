package schedules

import (
	"context"
	"fmt"
	"github.com/hsmade/velero-ui/internal/k8s"
	velerov1api "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

type ScheduleDetail struct {
	Schedule *velerov1api.Schedule `json:"schedule"`
	Backups  []*velerov1api.Backup `json:"backups"`
}

// GetScheduleDetail returns a struct with the schedule info for the schedule specified by `name`
func GetScheduleDetail(ctx context.Context, dynamicClient dynamic.Interface, name string) (*ScheduleDetail, error) {
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

// getSchedule gets a single schedule by `name`
func getSchedule(ctx context.Context, dynamicClient dynamic.Interface, name string) (*velerov1api.Schedule, error) {
	schedule := new(velerov1api.Schedule)
	err := k8s.GetByName(ctx, dynamicClient, schema.GroupVersionResource{
		Group:    "velero.io",
		Version:  "v1",
		Resource: "schedules",
	},
		name,
		&schedule,
	)
	return schedule, err
}

// getBackups gets all Backups by schedule `name`
func getBackups(ctx context.Context, dynamicClient dynamic.Interface, name string) ([]*velerov1api.Backup, error) {
	backups := new([]*velerov1api.Backup)
	err := k8s.ListByFilter(ctx, dynamicClient,
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
