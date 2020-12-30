package k8s

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetPodsByLabel (k8sClient *kubernetes.Clientset, namespace, label string) (*v1.PodList, error){
	return k8sClient.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{LabelSelector: label})
}

func GetPodsByName (k8sClient *kubernetes.Clientset, namespace, name string) (*v1.PodList, error){
	label := fmt.Sprintf("app=%s", name)
	return k8sClient.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{LabelSelector: label})
}