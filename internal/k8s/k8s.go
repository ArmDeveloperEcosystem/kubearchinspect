package k8s

import (
	"context"
	"os"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type KubernetesClient struct {
	clientset kubernetes.Clientset
}

func NewKubernetesClient() (*KubernetesClient, error) {
	// TODO: allow user to provide their own kubeconfig location
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	kubeconfig := filepath.Join(home, ".kube", "config")

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &KubernetesClient{
		clientset: *clientset,
	}, nil
}

func (k *KubernetesClient) GetAllPods() ([]corev1.Pod, error) {
	pods, err := k.clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return pods.Items, nil
}

func (k *KubernetesClient) GetAllImages() ([]string, error) {
	pods, err := k.GetAllPods()
	if err != nil {
		return nil, err
	}

	imageSet := make(map[string]struct{})
	for _, pod := range pods {
		for _, container := range pod.Spec.InitContainers {
			if _, found := imageSet[container.Image]; !found {
				imageSet[container.Image] = struct{}{}
			}
		}
		for _, container := range pod.Spec.Containers {
			if _, found := imageSet[container.Image]; !found {
				imageSet[container.Image] = struct{}{}
			}
		}
	}

	images := make([]string, 0, len(imageSet))
	for image := range imageSet {
		images = append(images, image)
	}
	return images, nil
}
