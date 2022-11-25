package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 加载kubeconfig文件，生成config对象
	config, err := clientcmd.BuildConfigFromFlags("", "C:\\Users\\hanwei\\config")

	if err != nil {
		panic(err)
	}
	// 配置API路径和请求的资源组/资源版本信息
	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	// 通过rest.RESTClientFor()生成RESTClient对象。 RESTClientFor通过令牌桶算法，有限制的说法。
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	// 通过RESTClient构建请求参数，查询default空间下所有pod资源
	result := &corev1.PodList{}
	err = restClient.Get().
		Namespace("default").
		Resource("pods").
		VersionedParams(&metav1.ListOptions{Limit: 500}, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(result)

	if err != nil {
		panic(err)
	}

	for _, d := range result.Items {
		fmt.Printf("NAMESPACE:%v \t NAME: %v \t STATUS: %v\n", d.Namespace, d.Name, d.Status.Phase)
	}
}
