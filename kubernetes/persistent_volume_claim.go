package kubernetes

import (
	"fmt"

	kube "k8s.io/client-go/kubernetes"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)
type PersistentVolumeClaim struct {
	ClientSet *kube.Clientset
}

func (p *PersistentVolumeClaim) DeleteByRelease(releaseName string, namespace string) error {
	deleteOptions := &meta_v1.DeleteOptions{
	}
	listOptions := meta_v1.ListOptions{
		LabelSelector: fmt.Sprintf("release=%s", releaseName),
	}
	return p.ClientSet.CoreV1().PersistentVolumeClaims(namespace).DeleteCollection(deleteOptions, listOptions)
}
