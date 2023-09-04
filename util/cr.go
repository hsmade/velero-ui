package util

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

func GetByName(ctx context.Context, dynamicClient dynamic.Interface, gvr schema.GroupVersionResource, name string, result interface{}) error {
	crRaw, err := dynamicClient.
		Resource(gvr).
		Namespace("velero").
		Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("fetching backup resource: %w", err)
	}

	err = runtime.DefaultUnstructuredConverter.FromUnstructured(crRaw.UnstructuredContent(), result)
	if err != nil {
		return fmt.Errorf("unmarshal backup resource: %w", err)
	}
	return nil
}

func GetByFilter(ctx context.Context, dynamicClient dynamic.Interface, gvr schema.GroupVersionResource, filter metav1.ListOptions, result interface{}) error {
	crRaw, err := dynamicClient.
		Resource(gvr).
		Namespace("velero").
		List(ctx, filter)
	if err != nil {
		return fmt.Errorf("listing backup deletion request resource: %w", err)
	}

	if len(crRaw.Items) == 0 {
		return nil
	}

	err = runtime.DefaultUnstructuredConverter.FromUnstructured(crRaw.Items[0].UnstructuredContent(), result)
	if err != nil {
		return fmt.Errorf("unmarshal backup deletion request resource: %w", err)
	}

	if len(crRaw.Items) > 1 {
		return fmt.Errorf("more than one backup deletion request returned: %d", len(crRaw.Items))
	}
	return nil
}

func ListByFilter[E interface{}](ctx context.Context, dynamicClient dynamic.Interface, gvr schema.GroupVersionResource, filter metav1.ListOptions, result *[]E) error {
	crRaw, err := dynamicClient.
		Resource(gvr).
		Namespace("velero").
		List(ctx, filter)
	if err != nil {
		return fmt.Errorf("listing backup deletion request resource: %w", err)
	}

	if len(crRaw.Items) == 0 {
		return nil
	}

	for _, itemRaw := range crRaw.Items {
		var item E
		err = runtime.DefaultUnstructuredConverter.FromUnstructured(itemRaw.UnstructuredContent(), item)
		if err != nil {
			return fmt.Errorf("unmarshal backup deletion request resource: %w", err)
		}
		*result = append(*result, item)
	}
	return nil
}
