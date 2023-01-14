package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var MG_APP_PORT = "30002"

type ClusterUsageRate struct {
	CPU    float64
	Memory float64
}
type NodeUsageRateResponse struct {
	CpuUsageRate []float64 `json:"cpu_usage_rate"`
}

func getNodeUsageRate(nodeIP string) (*NodeUsageRateResponse, error) {
	url := "http://" + nodeIP + ":" + MG_APP_PORT + "/cpu_usage_rate"
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var nodeUsageRateResponse NodeUsageRateResponse

	if err = json.Unmarshal(body, &nodeUsageRateResponse); err != nil {
		return nil, err
	}

	return &nodeUsageRateResponse, nil
}

func getClusterUsageRate(clientset *kubernetes.Clientset) (float64, error) {
	node, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return 0.0, err
	}

	var cpuList []float64
	for i := 0; i < len(node.Items); i++ {
		// master nodeは除外
		masterFlag := node.Items[i].Labels["node-role.kubernetes.io/control-plane"]
		if masterFlag == "true" {
			continue
		}
		// NodeのIPアドレスを取得
		nodeIP := node.Items[i].Status.Addresses[0].Address
		// NodeのCPU使用率を取得
		response, _ := getNodeUsageRate(nodeIP)

		if err != nil {
			return 0.0, err
		}
		// CPU使用率をリストに追加
		cpuList = append(cpuList, response.CpuUsageRate...)
	}

	// CPU使用率の平均を計算
	sum := 0.0
	for _, value := range cpuList {
		sum += value
	}
	cpuAverage := float64(sum) / float64(len(cpuList))
	return cpuAverage, nil
}

func setupK8sClient() (*kubernetes.Clientset, error) {
	var config *rest.Config
	var err error
	kubeconfig := "./k3s.yaml"
	config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

func getClusterUsageRateHandler(ctx *gin.Context) {
	clientset, err := setupK8sClient()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	clusterCpuUsageRate, err := getClusterUsageRate(clientset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"cpu": clusterCpuUsageRate,
	})
}
