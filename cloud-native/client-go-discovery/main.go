package main

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 加载kubeconfig文件，生成config对象
	config, err := clientcmd.BuildConfigFromFlags("", "C:\\Users\\hanwei\\config")
	if err != nil {
		panic(err)
	}

	// discovery.NewDiscoveryClientForConfigg函数通过config实例化discoveryClient对象
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		panic(err)
	}

	// discoveryClient.ServerGroupsAndResources 返回API Server所支持的资源组、资源版本、资源信息
	_, APIResourceList, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		panic(err)
	}

	// 输出所有资源信息
	for _, list := range APIResourceList {
		gv, err := schema.ParseGroupVersion(list.GroupVersion)
		if err != nil {
			panic(err)
		}

		for _, resource := range list.APIResources {
			fmt.Printf("NAME: %v, GROUP: %v, VERSION: %v \n", resource.Name, gv.Group, gv.Version)
		}
	}
}
