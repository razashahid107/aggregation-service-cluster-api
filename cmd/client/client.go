package client

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// default method for generating config
func GenerateDefaultConfig() *rest.Config {
	kubeconfig := flag.String("kubeconfig", "/Users/abdurrehman/.kube/config", "location for my kubeconfig file")
	// create config object 
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println("Config error")
	}
	return config
}

// generate config using context name
func BuildConfigWithContextFromFlags(context string, kubeconfigPath string) (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath},
		&clientcmd.ConfigOverrides{
			CurrentContext: context,
		}).ClientConfig()
}
// generate client using provided config
func Client(config *rest.Config) *kubernetes.Clientset {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("clientset error")
	}
	return clientset
}


func Generate_client() {
	kubeconfig := flag.String("kubeconfig", "/Users/abdurrehman/.kube/config", "location for my kubeconfig file")
	// create config object to create Kubernetes clients down the line
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println("Config error")
	}
	// create a Kubernetes client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("clientset error")
	}
	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("pods error")
	}
	fmt.Println("Pods: ")
	for _, pod := range pods.Items {
		fmt.Printf("%s", pod.Name)
	}
	node, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	// node, err := clientset.CoreV1().Nodes()
	if err != nil {
		fmt.Println("nodes could not get fetched")
	}
	fmt.Println()
	for _, nodes := range node.Items {
		fmt.Println("Nodes: ", nodes.Name)
		fmt.Println("Status: ", nodes.Status.Conditions)
	}
	// cluster, err := clientset.CoreV1().Clusters("default").List(context.TODO(), metav1.ListOptions{})
	// if err != nil {
	// 	fmt.Println("cluster not found error")
	// }
	// fmt.Println()
	// for _, clusters := range cluster.Items {

	// }

	// run the function existing in api.go
}
