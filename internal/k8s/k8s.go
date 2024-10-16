/*
Copyright 2024 Arm Limited

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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

// GetAllPods returns all Pods in all namespaces in the cluster
func (k *KubernetesClient) GetAllPods() ([]corev1.Pod, error) {
	pods, err := k.clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return pods.Items, nil
}

// GetAllImages returns all unique images used by all current running Pods in the cluster
// TODO: Get images from Deployments, CronJobs, etc which may not be running.
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
