package main

import (
	"fmt"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"ingress-monitor/structures"
)

var array []structures.ClientIngress

func GetIngress(namespaceList []string, clientset *kubernetes.Clientset) []structures.ClientIngress{

	for _, ns := range namespaceList {

		ingressList, err := clientset.NetworkingV1().Ingresses(ns).List(context.TODO() ,metav1.ListOptions{})

		if err != nil {
			fmt.Println(err)
		}

		for _, ingress := range ingressList.Items {
			newIngress := structures.ClientIngress{
				Name: ingress.Name,
				Host: ingress.Spec.Rules[0].Host,
				Namespace: ingress.Namespace,
			}

			array = append(array, newIngress)
		}

	}

	return array

}