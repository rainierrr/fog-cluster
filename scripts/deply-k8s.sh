#!/bin/bash
# Deploy the Kubernetes manifests to all cluster
kubectl apply -k ./kubernetes/ --kubeconfig=$HOME/.kube/k3s-cluster-a.yaml
kubectl apply -k ./kubernetes/ --kubeconfig=$HOME/.kube/k3s-cluster-b.yaml
