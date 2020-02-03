#!/bin/sh
kubectl create -f provision/k8s/deployment.yaml

# # get deployment
# kubectl get deployments

# # get service
# kubectl get services

# # delete service
# kubectl delete service tranngocdan-nc_student
# kubectl delete service nc_student

# # get pods
# kubectl get pods

# minikube service tranngocdan-nc_student-service --url