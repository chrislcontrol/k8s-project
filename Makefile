run:
	- docker run --name go-k8s --rm -p 8000:8000 go-k8s

create-cluster-kind:
	- kind create cluster

use-context:
	- kubectl config use-context kind-kind

apply config:
	- kubectl apply -f k8s/deployment

delete-cluster-kind:
	- kind delete cluster
