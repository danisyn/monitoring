package main

import (
	"fmt"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetIngress(ns string, clientset *kubernetes.Clientset) []ClientIngress{

	ingressList, err := clientset.NetworkingV1().Ingresses(ns).List(context.TODO() ,metav1.ListOptions{})

	if err != nil {
		fmt.Println(err)
	}

	for _, ingress := range ingressList.Items {
		newIngress := ClientIngress{
			Name: ingress.Name,
			Host: ingress.Spec.Rules[0].Host,
		}

		array = append(array, newIngress)
	}

	return array

}