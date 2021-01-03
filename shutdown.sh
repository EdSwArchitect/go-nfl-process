#!/bin/bash

minikube kubectl -- delete po kafka01 nfl-processor read-nfl-processor zookeeper

minikube kubectl -- delete svc kafka-svc kafka01 nfl-processor zookeeper

