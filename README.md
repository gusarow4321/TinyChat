# TinyChat

![CI/CD](https://github.com/gusarow4321/TinyChat/workflows/CI/CD/badge.svg)
[![Codecov](https://img.shields.io/codecov/c/github/gusarow4321/TinyChat)](https://codecov.io/gh/gusarow4321/TinyChat)

- [x] Microservices framework: [go-kratos](https://github.com/go-kratos/kratos)
- [x] Api Gateway: [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
- [x] OpenAPI descriptions with [protoc-gen-openapi](https://github.com/google/gnostic/tree/main/cmd/protoc-gen-openapi)
- [x] PostgreSQL database
- [x] Unit tests with [testify](https://github.com/stretchr/testify)
- [x] Tracing: [Jaeger](https://www.jaegertracing.io/)
- [x] Metrics: [Prometheus](https://prometheus.io/) & [Grafana](https://grafana.com/) dashboards
- [x] Messages Broker: Kafka
- [x] Orchestration: [K8s](https://kubernetes.io/)
- [ ] File storage
- [ ] Deploying to GKE
- [ ] Integration tests

## Build
Build docker images of all services:
```shell
make all-imgs
```

## Docker
Deploy services locally with docker:
```shell
docker compose up -d
```

## Kubernetes
Deploy services on [minikube](https://minikube.sigs.k8s.io/docs/start/) using [helm](https://helm.sh/)
```shell
minikube start --feature-gates=GRPCContainerProbe=true

kubectl create namespace tiny-chat
kubectl config set-context --current --namespace=tiny-chat

kubectl create secret generic grafana-auth \
  --from-literal=admin-user=admin \
  --from-literal=admin-password=admin

make load-imgs

helm dependency update
helm install tiny-chat tiny-chat-chart
```

To forward pod's port to localhost:
```shell
kubectl port-forward POD_NAME LOCALHOST_PORT:POD_PORT
```
