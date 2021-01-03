#!/bin/bash

minikube kubectl -- apply -f kafka-0mc-pods.yaml
minikube kubectl -- apply -f nfl-processor.yaml
minikube kubectl -- apply -f read-nfl-processor.yaml

echo "List pods"

minikube kubectl -- get po -o wide

echo "List services"

minikube kubectl -- get svc -o wide

sleep 3m

minikube kubectl -- apply -f nfl-processor.yaml

