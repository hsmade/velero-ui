package backups

import (
	"context"
	"fmt"
	velerov1api "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
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

	podvolumebackups, err := getPodVolumeBackups(ctx, dynamicClient, backup.UID)
	if err != nil {
		return nil, fmt.Errorf("get podvolume backups: %w", err)
	}

	result := BackupDetail{
		Backup:              backup,
		DeleteBackupRequest: deletionRequest,
		PodVolumeBackups:    podvolumebackups,
	}

	return &result, nil
}

func getBackup(ctx context.Context, dynamicClient dynamic.Interface, name string) (*velerov1api.Backup, error) {
	backupRaw, err := dynamicClient.
		Resource(schema.GroupVersionResource{
			Group:    "velero.io",
			Version:  "v1",
			Resource: "backups",
		}).
		Namespace("velero").
		Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("fetching backup resource: %w", err)
	}

	backup := new(velerov1api.Backup)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(backupRaw.UnstructuredContent(), &backup)
	if err != nil {
		return nil, fmt.Errorf("unmarshal backup resource: %w", err)
	}

	return backup, nil
}

func getBackupDeleteRequest(ctx context.Context, dynamicClient dynamic.Interface, uid types.UID) (*velerov1api.DeleteBackupRequest, error) {
	result, err := dynamicClient.
		Resource(schema.GroupVersionResource{
			Group:    "velero.io",
			Version:  "v1",
			Resource: "deletebackuprequests",
		}).
		Namespace("velero").
		List(ctx, metav1.ListOptions{LabelSelector: fmt.Sprintf("velero.io/backup-uid=%s", uid)})
	if err != nil {
		return nil, fmt.Errorf("listing backup deletion request resource: %w", err)
	}

	if len(result.Items) == 0 {
		return nil, nil
	}

	request := new(velerov1api.DeleteBackupRequest)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(result.Items[0].UnstructuredContent(), &request)
	if err != nil {
		return nil, fmt.Errorf("unmarshal backup deletion request resource: %w", err)
	}

	if len(result.Items) > 1 {
		return request, fmt.Errorf("more than one backup deletion request returned: %d", len(result.Items))
	}
	return request, nil
}

func getPodVolumeBackups(ctx context.Context, dynamicClient dynamic.Interface, uid types.UID) ([]*velerov1api.PodVolumeBackup, error) {
	result, err := dynamicClient.
		Resource(schema.GroupVersionResource{
			Group:    "velero.io",
			Version:  "v1",
			Resource: "podvolumebackups",
		}).
		Namespace("velero").
		List(ctx, metav1.ListOptions{LabelSelector: fmt.Sprintf("velero.io/backup-uid=%s", uid)})
	if err != nil {
		return nil, fmt.Errorf("listing pod volume backup resource: %w", err)
	}

	if len(result.Items) == 0 {
		return nil, nil
	}

	backups := []*velerov1api.PodVolumeBackup{}
	for _, item := range result.Items {
		backup := new(velerov1api.PodVolumeBackup)
		err = runtime.DefaultUnstructuredConverter.FromUnstructured(item.UnstructuredContent(), &backup)
		if err != nil {
			return nil, fmt.Errorf("unmarshal podvolume backup resource: %w", err)
		}
		backups = append(backups, backup)
	}

	return backups, nil
}
