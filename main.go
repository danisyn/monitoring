package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"ingress-monitor/structures"
	"log"
	"net/http"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func router() {
	router := http.NewServeMux()
	router.HandleFunc("/", writeIngress)

	log.Print("Listening...")
    err := http.ListenAndServe(":3001", router)
    if err != nil {
        log.Fatal(err)
    }
}

func writeIngress(w http.ResponseWriter, r *http.Request) {
	ingress := firstFunc()
	var response bytes.Buffer
	enc := gob.NewEncoder(&response)
	enc.Encode(ingress)

	w.Write(response.Bytes())
}

func firstFunc() []structures.ClientIngress{
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

	hosts := GetIngress(ns, clientset)

	return hosts
}

func namespaces(clientset *kubernetes.Clientset) []string {
	var array []string

	nsList, _ := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})


	for _, v := range nsList.Items {
		array = append(array, v.Name)
	}

	return array
}

func main() {
	router()
}