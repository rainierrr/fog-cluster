package main

import (
	"context"
	"errors"
	"fmt"
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

	clusterCpuAllocatable := float64(0)
	clusterMemAllocatable := float64(0)
	clusterCpuUsed := float64(0)
	clusterMemUsed := float64(0)

	if len(nodeMetricses.Items) == 0 {
		return nil, errors.New("nodeMetricses.Items is empty")
	}

	for i := 0; i < len(nodeMetricses.Items); i++ {
		// master nodeは除外
		masterFlag := node.Items[i].Labels["node-role.kubernetes.io/control-plane"]
		if masterFlag == "true" {
			continue
		}
		cpuUsage := float64(nodeMetricses.Items[i].Usage.Cpu().MilliValue())
		memUsage := float64(nodeMetricses.Items[i].Usage.Memory().MilliValue())
		cpuAllocatable := float64(node.Items[i].Status.Allocatable.Cpu().MilliValue())
		memAllocatable := float64(node.Items[i].Status.Allocatable.Memory().MilliValue())

		clusterCpuAllocatable += cpuAllocatable
		clusterMemAllocatable += memAllocatable
		clusterCpuUsed += cpuUsage
		clusterMemUsed += memUsage

		//デバック用
		fmt.Printf("Node name: %s\n", node.Items[i].Name)
		fmt.Printf("CPU rate: %f\n", cpuUsage/cpuAllocatable)
		fmt.Printf("Memory rate: %f\n", memUsage/memAllocatable)
		fmt.Printf("cpuUsage: %d\n", nodeMetricses.Items[i].Usage.Cpu().MilliValue())
		memUsage2 := node.Items[i].Status.Capacity.Memory().MilliValue()
		fmt.Printf("memUsage: %d\n", memUsage2)
		fmt.Println("--------------------------------------------------")
	}

	clusterUsageRate := ClusterUsageRate{
		CPU:    clusterCpuUsed / clusterCpuAllocatable,
		Memory: clusterMemUsed / clusterMemAllocatable,
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

	resouceUsageRate, err := getClusterUsageRate(clientset, mc)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"cpu": resouceUsageRate.CPU,
		// "memory": resouceUsageRate.Memory,
	})
}
