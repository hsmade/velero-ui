// Package k8s provides helper functions to query for the velero custom resources
package k8s

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"reflect"
)

// GetByName gets a single custom resource by Name
func GetByName(ctx context.Context, dynamicClient dynamic.Interface, gvr schema.GroupVersionResource, name string, result interface{}) error {
	crRaw, err := dynamicClient.
		Resource(gvr).
		Namespace("velero").
		Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("fetching '%s' resource: %w", reflect.TypeOf(result), err)
	}

	err = runtime.DefaultUnstructuredConverter.FromUnstructured(crRaw.UnstructuredContent(), result)
	if err != nil {
		return fmt.Errorf("unmarshal '%s' resource: %w", reflect.TypeOf(result), err)
	}
	return nil
}

// GetByFilter gets a single custom resource by filter (labels, etc)
func GetByFilter(ctx context.Context, dynamicClient dynamic.Interface, gvr schema.GroupVersionResource, filter metav1.ListOptions, result interface{}) error {
	crRaw, err := dynamicClient.
		Resource(gvr).
		Namespace("velero").
		List(ctx, filter)
	if err != nil {
		return fmt.Errorf("listing '%s' resource: %w", reflect.TypeOf(result), err)
	}

	if len(crRaw.Items) == 0 {
		return nil
	}

	err = runtime.DefaultUnstructuredConverter.FromUnstructured(crRaw.Items[0].UnstructuredContent(), result)
	if err != nil {
		return fmt.Errorf("unmarshal '%s'' resource: %w", reflect.TypeOf(result), err)
	}

	if len(crRaw.Items) > 1 {
		return fmt.Errorf("more than one '%s' returned: %d", reflect.TypeOf(result), len(crRaw.Items))
	}
	return nil
}

// ListByFilter gets all custom resources that match a filter (labels, etc)
func ListByFilter[E interface{}](ctx context.Context, dynamicClient dynamic.Interface, gvr schema.GroupVersionResource, filter metav1.ListOptions, result *[]E) error {
	crRaw, err := dynamicClient.
		Resource(gvr).
		Namespace("velero").
		List(ctx, filter)
	if err != nil {
		return fmt.Errorf("listing '%s'' resource: %w", reflect.TypeOf(result), err)
	}

	if len(crRaw.Items) == 0 {
		return nil
	}

	for _, itemRaw := range crRaw.Items {
		var item E
		err = runtime.DefaultUnstructuredConverter.FromUnstructured(itemRaw.UnstructuredContent(), &item)
		if err != nil {
			return fmt.Errorf("unmarshal '%s' resource: %w", reflect.TypeOf(result), err)
		}
		*result = append(*result, item)
	}
	return nil
}
