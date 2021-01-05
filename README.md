# Minikube and Kafka 

## Mounting local drive with Minikube
minikube mount $HOME:/host

## Running

1. minikube kubectl -- apply -f kafka-0mc-pods.xml
1. minikube kubectl -- apply -f nfl-processor.yaml
1. minikube kubectl -- apply -f read-nfl-processor.yaml

### For Config Map processing

1. minikube kubectl -- create configmap resources --file-file=./resources 
1. minikube kubectl -- apply -f nfl-processor-config.yaml
1. minikube kubectl -- apply -f read-nfl-processor.yaml

## Cleanup

1. minikube kubectl -- delete po kafka01 nfl-processor read-nfl-processor zookeeper
1. minikube kubectl -- delete svc kafka-svc kafka01 nfl-processor zookeeper

