package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
	"regexp"
)

func main() {
	_, inCluster := os.LookupEnv("KUBERNETES_SERVICE_HOST")

	var config *rest.Config
	var err error

	if inCluster {
		config, err = rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
	} else {
		fmt.Println("[ERROR]: This code requires to be in a kubernetes pod")
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	ns := namespaces(clientset)

	GetIngress(ns, clientset)

}

func namespaces(clientset *kubernetes.Clientset) string {
	var array string

	nsList, _ := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})

	filter := "-prd"
	reg, _ := regexp.Compile(filter)

	for _, v := range nsList.Items {
		match := reg.MatchString(v.Name)

		if match {
			array = v.Name
		}
	}

	return array
}