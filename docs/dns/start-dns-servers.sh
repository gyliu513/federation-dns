#!/usr/bin/env bash

base_dir="$( cd "$(dirname "$0")/../.." && pwd )"
base_dir=${base_dir##$(pwd)/}

source ${base_dir}/docs/common/util.sh

if [[ $# -lt 1 ]]; then
  echo "usage: $0  <cluster-name>  <variant(v1|v2)>"
  exit 1
fi

Cluster=$1
Variant="v1"
if [[ -n "$2" ]]; then
  Variant=$2
fi
NS="fed-dns"

read -s
clear

function wait_for_pods_to_be_ready() {
  Deployment=$1
  while [ "0" == "$(kubectl -n ${NS} get deployments ${Deployment} -o jsonpath="{.status.availableReplicas}")" ]; do
    echo "Waiting for ${Deployment} pod to be Running"
    sleep 3;
  done
  run "kubectl -n ${NS} get pods"
}

#DNS Controller
kubectl config use-context ${Cluster}

run "kubectl create ns ${NS}"

run "kubectl -n ${NS} run etcd --image=quay.io/coreos/etcd:v3.3 --env="ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:2379" --env="ETCD_ADVERTISE_CLIENT_URLS=http://etcd.${NS}:2379" --port=2379 --expose"
wait_for_pods_to_be_ready "etcd"

run "helm install --namespace ${NS} --name andromeda -f ${base_dir}/docs/dns/config/coredns-chart-values.yaml stable/coredns"
wait_for_pods_to_be_ready "andromeda-coredns"

run "helm install --namespace ${NS} --name milkyway -f ${base_dir}/docs/dns/config/federation-dns-chart-values.yaml ${base_dir}/chart/federation-dns --set variant=${Variant}"
wait_for_pods_to_be_ready "milkyway-federation-dns"

if [[ "${Variant}" == "v2" ]]; then
  run "helm install --namespace ${NS} --name whirlpool -f ${base_dir}/docs/dns/config/external-dns-chart-values.yaml stable/external-dns"
  wait_for_pods_to_be_ready "whirlpool-external-dns"
fi

# TODO: Obtain the ip of CoreDNS service and configure the kube-dns-config, instead of hard coding as is done now.
for filename in $(ls ${base_dir}/docs/dns/config/kube-dns-configmap*.yaml); do
  kubectl create -f "${filename}"
done

run "kubectl -n ${NS} get pods"
