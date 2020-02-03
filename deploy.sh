#!/bin/sh
kubectl delete service tranngocdan-nc-student-service
kubectl delete deployment tranngocdan-nc-student
kubectl create -f provision/k8s/deployment.yaml
kubectl get service tranngocdan-nc-student-service

# minikube service tranngocdan-nc-student-service --url