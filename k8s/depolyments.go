package k8s

import (
	"context"
	"fmt"
	"strings"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetDps(k8sClient *kubernetes.Clientset, namespace string) (*v1.DeploymentList, error) {
	return k8sClient.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
}

func GetDpByName(k8sClient *kubernetes.Clientset, name string, namespace string) (*v1.Deployment, error) {
	return k8sClient.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
}

func GetImageByDpName(k8sClient *kubernetes.Clientset, name, namespace string) (imageName, version string, err error) {
	dp, err := GetDpByName(k8sClient, name, namespace)
	if err != nil {
		return "", "", err
	}
	for _, ctr := range dp.Spec.Template.Spec.Containers {
		if strings.Contains(ctr.Name, dp.GetName()) {
			for _, en := range ctr.Env {
				if en.Name == "VERSION" {
					version = en.Value
					break
				}
			}
			return ctr.Image, version, nil
		}
	}
	return "", "", fmt.Errorf("unable to find image list")
}

func SetDpImage(k8sClient *kubernetes.Clientset, name, namespace, image string) (err error) {
	dp, err := GetDpByName(k8sClient, name, namespace)
	if err != nil {
		return err
	}
	found := false
	for i, ctr := range dp.Spec.Template.Spec.Containers {
		if strings.Contains(ctr.Name, dp.GetName()) {
			found = true
			dp.Spec.Template.Spec.Containers[i].Image = image
			break
		}
	}
	if !found {
		return fmt.Errorf("the application container not exist in the deployment pods")
	}
	_, err = k8sClient.AppsV1().Deployments(namespace).Update(context.TODO(), dp, metav1.UpdateOptions{})
	return err
}
