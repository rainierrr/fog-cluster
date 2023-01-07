package main

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

type ClusterUsageRate struct {
	CPU    float64
	Memory float64
}

func getClusterUsageRate(clientset *kubernetes.Clientset, mc *metrics.Clientset) (*ClusterUsageRate, error) {
	nodeMetricses, err := mc.MetricsV1beta1().NodeMetricses().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	node, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	if len(nodeMetricses.Items) != len(node.Items) {
		return nil, errors.New("nodeMetricses.Items and node.Items are not equal")
	}
	cluster_cpu_allocatable := float64(0)
	cluster_mem_allocatable := float64(0)
	cluster_mem_used := float64(0)
	cluster_cpu_used := float64(0)

	for i := 0; i < len(nodeMetricses.Items); i++ {
		cpu_usage := float64(nodeMetricses.Items[i].Usage.Cpu().MilliValue())
		mem_usage := float64(nodeMetricses.Items[i].Usage.Memory().Value())
		cpu_allocatable := float64(node.Items[i].Status.Allocatable.Cpu().MilliValue())
		mem_allocatable := float64(node.Items[i].Status.Allocatable.Memory().Value())

		cluster_cpu_allocatable += cpu_allocatable
		cluster_mem_allocatable += mem_allocatable
		cluster_cpu_used += cpu_usage
		cluster_mem_used += mem_usage

		// デバック用
		// fmt.Printf("Node name: %s\n", node.Items[i].Name)
		// fmt.Printf("CPU rate: %f\n", cpu_usage/cpu_allocatable)
		// fmt.Printf("Memory rate: %f\n", mem_usage/mem_allocatable)
	}

	clusterUsageRate := ClusterUsageRate{
		CPU:    cluster_cpu_used / cluster_cpu_allocatable,
		Memory: cluster_mem_used / cluster_mem_allocatable,
	}

	return &clusterUsageRate, nil
}

func setupK8sClient() (*kubernetes.Clientset, *metrics.Clientset, error) {
	var config *rest.Config
	var err error
	if os.Getenv("ENV") == "local" {
		kubeconfig := "./k3s.yaml"
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	} else {
		config, err = rest.InClusterConfig()
	}

	if err != nil {
		return nil, nil, err
	}

	mc, err := metrics.NewForConfig(config)
	if err != nil {
		return nil, nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, nil, err
	}

	return clientset, mc, nil
}

func getClusterUsageRateHandler(ctx *gin.Context) {
	clientset, mc, err := setupK8sClient()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	resouce_usage_rate, err := getClusterUsageRate(clientset, mc)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"cpu":    resouce_usage_rate.CPU,
		"memory": resouce_usage_rate.Memory,
	})
}
