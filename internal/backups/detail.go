package backups

import (
	"context"
	"fmt"
	"github.com/hsmade/velero-ui/internal/k8s"
	velerov1api "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
)

type BackupDetail struct {
	Backup              *velerov1api.Backup              `json:"backup"`
	DeleteBackupRequest *velerov1api.DeleteBackupRequest `json:"delete_backup_request"`
	PodVolumeBackups    []*velerov1api.PodVolumeBackup   `json:"pod_volume_backups"`
}

// GetBackupDetail returns a struct with the backup info for the backup specified by `name`
func GetBackupDetail(ctx context.Context, dynamicClient dynamic.Interface, name string) (*BackupDetail, error) {
	backup, err := getBackup(ctx, dynamicClient, name)
	if err != nil {
		return nil, fmt.Errorf("get backup: %w", err)
	}

	deletionRequest, err := getBackupDeleteRequest(ctx, dynamicClient, backup.UID)
	if err != nil {
		return nil, fmt.Errorf("get backup delete request: %w", err)
	}

	podVolumeBackups, err := getPodVolumeBackups(ctx, dynamicClient, backup.UID)
	if err != nil {
		return nil, fmt.Errorf("get podvolume backups: %w", err)
	}

	result := BackupDetail{
		Backup:              backup,
		DeleteBackupRequest: deletionRequest,
		PodVolumeBackups:    podVolumeBackups,
	}

	return &result, nil
}

// getBackup gets a single backup by `name`
func getBackup(ctx context.Context, dynamicClient dynamic.Interface, name string) (*velerov1api.Backup, error) {
	backup := new(velerov1api.Backup)
	err := k8s.GetByName(ctx, dynamicClient, schema.GroupVersionResource{
		Group:    "velero.io",
		Version:  "v1",
		Resource: "backups",
	},
		name,
		&backup,
	)
	return backup, err
}

// getBackupDeleteRequest gets all BackupDeleteRequests by backup `uid`
func getBackupDeleteRequest(ctx context.Context, dynamicClient dynamic.Interface, uid types.UID) (*velerov1api.DeleteBackupRequest, error) {
	backupDeleteRequest := new(velerov1api.DeleteBackupRequest)
	err := k8s.GetByFilter(ctx, dynamicClient, schema.GroupVersionResource{
		Group:    "velero.io",
		Version:  "v1",
		Resource: "deletebackuprequests",
	},
		metav1.ListOptions{LabelSelector: fmt.Sprintf("velero.io/backup-uid=%s", uid)},
		&backupDeleteRequest,
	)
	return backupDeleteRequest, err
}

// getPodVolumeBackups gets all PodVolumeBackups by backup `uid`
func getPodVolumeBackups(ctx context.Context, dynamicClient dynamic.Interface, uid types.UID) ([]*velerov1api.PodVolumeBackup, error) {
	podVolumeBackups := new([]*velerov1api.PodVolumeBackup)
	err := k8s.ListByFilter(ctx, dynamicClient,
		schema.GroupVersionResource{
			Group:    "velero.io",
			Version:  "v1",
			Resource: "podvolumebackups",
		},
		metav1.ListOptions{LabelSelector: fmt.Sprintf("velero.io/backup-uid=%s", uid)},
		podVolumeBackups,
	)
	return *podVolumeBackups, err
}
