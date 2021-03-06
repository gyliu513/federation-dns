#!/usr/bin/env bash

base_dir="$( cd "$(dirname "$0")/../.." && pwd )"
base_dir=${base_dir##$(pwd)/}

source ${base_dir}/docs/common/util.sh

if [[ $# -ne 3 ]]; then
  echo "usage: $0  <cluster-name>  <federation-member-1>  <federation-member-2>"
  exit 1
fi

Cluster=$1
C1=$2
C2=$3

read -s
clear

# Bringup k8s cluster, followed by federation control plane in the cluster and join given 2 clusters to federation
kubectl config use-context ${Cluster}

while [ "Running" != "$(kubectl -n kube-system get pod storage-provisioner -o jsonpath="{.status.phase}")" ]; do
    echo "Waiting for storage-provisioner pod to be Running"
    sleep 3;
done
run "kubectl -n kube-system get pod storage-provisioner"
run "kubectl get all --all-namespaces"

run "crinit aggregated init mycr --host-cluster-context=${Cluster}"

run "kubectl create ns federation"

run "kubectl apply -f ${base_dir}/docs/bringup/config/apiserver.yaml"

run "kubectl api-versions"

kubectl get federatedcluster
while [ $? -ne 0 ]; do
    sleep 3
    kubectl get federatedcluster
done
run "kubectl get federatedcluster"

# Enable replicaset and service types
for filename in ${base_dir}/docs/bringup/config/federatedtypes/*.yaml; do
  kubectl create -f "${filename}"
done

run "kubefnord join ${C1} --host-cluster-context ${Cluster} --add-to-registry --v=2"
run "kubefnord join ${C2} --host-cluster-context ${Cluster} --add-to-registry --v=2"

run "kubectl get federatedcluster"
