package main

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 加载kubeconfig文件，生成config对象
	config, err := clientcmd.BuildConfigFromFlags("", "C:\\Users\\hanwei\\config")
	if err != nil {
		panic(err)
	}

	// kubernetes.NewForConfig通过config实例化ClientSet对象
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	//请求core核心资源组v1资源版本下的Pods资源对象
	podClient := clientset.CoreV1().Pods(apiv1.NamespaceDefault)
	// 设置选项
	list, err := podClient.List(context.TODO(), metav1.ListOptions{Limit: 500})
	if err != nil {
		panic(err)
	}

	for _, d := range list.Items {
		fmt.Printf("NAMESPACE: %v \t NAME:%v \t STATUS: %+v\n", d.Namespace, d.Name, d.Status.Phase)
	}
}
