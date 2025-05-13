/*
Copyright 2025 Arm Limited

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
	"log"
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

func NewKubernetesClient(kubeconfigPath, context string, debug bool) (*KubernetesClient, error) {
	// Check if user provided the kubeconfig location
	if kubeconfigPath == "" {
		// Try the kubeconfig in the default location
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}

		kubeconfigPath = filepath.Join(home, ".kube", "config")
	}

	configOverrides := &clientcmd.ConfigOverrides{}
	if context != "" {
		configOverrides.CurrentContext = context
	}

	// Build the configuration from the kubeconfig file and overrides
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath},
		configOverrides,
	).ClientConfig()
	if err != nil {
		return nil, err
	}

	// Log the context being used
	if debug {
		log.Println("Using the kube config file at", kubeconfigPath)
		if context == "" {
			rawConfig, err := clientcmd.LoadFromFile(kubeconfigPath)
			if err != nil {
				return nil, err
			}
			context = rawConfig.CurrentContext
		}
		log.Println("Running the command against context:", context)
	}

	// Create the clientset
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
func (k *KubernetesClient) GetAllImages() (map[string][]string, error) {
	pods, err := k.GetAllPods()
	if err != nil {
		return nil, err
	}

	imageMap := make(map[string][]string)
	for _, pod := range pods {
		for _, container := range pod.Spec.InitContainers {
			imageMap[container.Image] = append(imageMap[container.Image], pod.Name)
		}
		for _, container := range pod.Spec.Containers {
			imageMap[container.Image] = append(imageMap[container.Image], pod.Name)
		}
	}

	return imageMap, nil
}
