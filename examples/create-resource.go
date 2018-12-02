// +build ignore

package configmaps

import (
	"context"

	"github.com/tryggth/k8s"
	"github.com/tryggth/k8s/apis/core/v1"
	metav1 "github.com/tryggth/k8s/apis/meta/v1"
)

func createConfigMap(client *k8s.Client, name string, values map[string]string) error {
	cm := &v1.ConfigMap{
		Metadata: &metav1.ObjectMeta{
			Name:      &name,
			Namespace: &client.Namespace,
		},
		Data: values,
	}
	// Will return the created configmap as well.
	return client.Create(context.TODO(), cm)
}
