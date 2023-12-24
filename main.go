package main

import (
	handlers "aggregation-service-cluster-api/cmd/api/handlers"
	client "aggregation-service-cluster-api/cmd/client"

	"github.com/gin-gonic/gin"
)

// import (
// 	"context"
// 	"flag"
// 	"fmt"

// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	"k8s.io/client-go/kubernetes"
// 	"k8s.io/client-go/tools/clientcmd"
// )

// func main() {
// 	kubeconfig := flag.String("kubeconfig", "/Users/abdurrehman/.kube/config", "location for my kubeconfig file")
// 	// create config object to create Kubernetes clients down the line
// 	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
// 	if err != nil {
// 		fmt.Println("Config error")
// 	}
// 	// create a Kubernetes client
// 	clientset, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		fmt.Println("clientset error")
// 	}
// 	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
// 	if err != nil {
// 		fmt.Println("pods error")
// 	}
// 	fmt.Println("Pods: ")
// 	for _, pod := range pods.Items {
// 		fmt.Printf("%s", pod.Name)
// 	}
// 	node, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
// 	// node, err := clientset.CoreV1().Nodes()
// 	if err != nil {
// 		fmt.Println("nodes could not get fetched")
// 	}
// 	fmt.Println()
// 	for _, nodes := range node.Items {
// 		fmt.Println("Nodes: ", nodes.Name)
// 		fmt.Println("Status: ", nodes.Status.Conditions)
// 	}
// 	// cluster, err := clientset.CoreV1().Clusters("default").List(context.TODO(), metav1.ListOptions{})
// 	// if err != nil {
// 	// 	fmt.Println("cluster not found error")
// 	// }
// 	// fmt.Println()
// 	// for _, clusters := range cluster.Items {

// 	// }

// 	// run the function existing in api.go
// 	run()
// }
func main() {
    router := gin.Default()
	// create a client 
	client := client.Client()
	// set nodes to handleListNodes function in cmd/api/handlers/list_nodes.go
	nodes, err := handlers.HandleListNodes(client)
	if err != nil {	
		panic(err)
	}
	pods, err := handlers.HandleListPods(client)
	if err != nil {	
		panic(err)
	}
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hey there! Go on to /nodes or /pods to get the list of nodes and pods respectively.",
		})
	})
    router.GET("/nodes", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"nodes": nodes,
		})
	})
	router.GET("/pods", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"nodes": pods,
		})
	})
    router.Run("localhost:8081")
}