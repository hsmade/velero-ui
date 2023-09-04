package backups

import (
	"context"
	"fmt"
	"github.com/hsmade/velero-ui/util"
	velerov1api "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
)

type BackupDetail struct {
	Backup              *velerov1api.Backup
	DeleteBackupRequest *velerov1api.DeleteBackupRequest
	PodVolumeBackups    []*velerov1api.PodVolumeBackup
}

func Detail(ctx context.Context, dynamicClient dynamic.Interface, name string) (*BackupDetail, error) {
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

func getBackup(ctx context.Context, dynamicClient dynamic.Interface, name string) (*velerov1api.Backup, error) {
	backup := new(velerov1api.Backup)
	err := util.GetByName(ctx, dynamicClient, schema.GroupVersionResource{
		Group:    "velero.io",
		Version:  "v1",
		Resource: "backups",
	},
		name,
		&backup,
	)
	return backup, err
}

func getBackupDeleteRequest(ctx context.Context, dynamicClient dynamic.Interface, uid types.UID) (*velerov1api.DeleteBackupRequest, error) {
	backupDeleteRequest := new(velerov1api.DeleteBackupRequest)
	err := util.GetByFilter(ctx, dynamicClient, schema.GroupVersionResource{
		Group:    "velero.io",
		Version:  "v1",
		Resource: "deletebackuprequests",
	},
		metav1.ListOptions{LabelSelector: fmt.Sprintf("velero.io/backup-uid=%s", uid)},
		&backupDeleteRequest,
	)
	return backupDeleteRequest, err
}

func getPodVolumeBackups(ctx context.Context, dynamicClient dynamic.Interface, uid types.UID) ([]*velerov1api.PodVolumeBackup, error) {
	podVolumeBackups := new([]*velerov1api.PodVolumeBackup)
	err := util.ListByFilter(ctx, dynamicClient,
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
