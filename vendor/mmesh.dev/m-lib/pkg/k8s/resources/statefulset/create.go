package statefulset

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/k8s/config"
)

func (a *API) Create(s *appsv1.StatefulSet) (*appsv1.StatefulSet, error) {
	if err := a.Delete(s.ObjectMeta.Namespace, s.ObjectMeta.Name); err != nil {
		return nil, errors.Wrapf(err, "[%v] function a.Get()", errors.Trace())
	}

	if a.clientset == nil {
		clientset, err := config.NewClient(a.KubeConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function config.NewClient()", errors.Trace())
		}
		a.clientset = clientset
	}

	ctx := context.TODO()

	createOpts := metav1.CreateOptions{}

	s, err := a.clientset.AppsV1().StatefulSets(s.ObjectMeta.Namespace).Create(ctx, s, createOpts)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function a.clientset.AppsV1().StatefulSets().Create()", errors.Trace())
	}

	return s, nil
}
