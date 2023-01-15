package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var MG_APP_PORT = "30002"
var MG_INTERNAL_PORT = "3000"

type ClusterUsageRate struct {
	CPU    float64
	Memory float64
}
type NodeUsageRateResponse struct {
	CpuUsageRate []float64 `json:"cpu_usage_rate"`
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getNodeUsageRate(podIP string) (*NodeUsageRateResponse, error) {
	url := "http://" + podIP + ":" + MG_INTERNAL_PORT + "/cpu_usage_rate"
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
	var cpuList []float64
	var masterNodeNameList []string
	// デーモンセットのPodを取得
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{
		LabelSelector: "node-role.kubernetes.io/master=true",
	})

	for _, node := range nodes.Items {
		masterNodeNameList = append(masterNodeNameList, node.Name)
	}
	if err != nil {
		return 0.0, err
	}
	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{
		LabelSelector: "app=mg-app",
	})
	if err != nil {
		return 0.0, err
	}

	for _, pod := range pods.Items {
		// master nodeに配置されたPodは除外
		if Contains(masterNodeNameList, pod.Spec.NodeName) {
			continue
		}
		podIP := pod.Status.PodIP

		// PodのCPU使用率を取得
		response, err := getNodeUsageRate(podIP)
		if err != nil {
			return 0.0, err
		}

		cpuList = append(cpuList, response.CpuUsageRate...)
	}

	// CPU使用率の平均を計算
	sum := 0.0
	for _, value := range cpuList {
		sum += value
	}
	cpuAverage := float64(sum) / float64(len(cpuList))
	log.Printf("CPU使用率の平均: %f\n", cpuAverage)
	log.Println("CPUのリスト", cpuList)
	return cpuAverage, nil
}

func setupK8sClient() (*kubernetes.Clientset, error) {
	var config *rest.Config
	var err error

	if os.Getenv("KUBERNETES_SERVICE_HOST") != "" {
		config, err = rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
	} else {
		kubeconfig := "./k3s.yaml"
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
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
