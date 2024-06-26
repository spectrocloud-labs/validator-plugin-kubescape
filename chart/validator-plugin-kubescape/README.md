
validator-plugin-kubescape
===========

validator-plugin-kubescape performs a variety of validations against your Azure account.


## Configuration

The following table lists the configurable parameters of the validator-plugin-kubescape chart and their default values.

| Parameter                | Description             | Default        |
| ------------------------ | ----------------------- | -------------- |
| `controllerManager.kubeRbacProxy.args` |  | `["--secure-listen-address=0.0.0.0:8443", "--upstream=http://127.0.0.1:8080/", "--logtostderr=true", "--v=0"]` |
| `controllerManager.kubeRbacProxy.containerSecurityContext.allowPrivilegeEscalation` |  | `false` |
| `controllerManager.kubeRbacProxy.containerSecurityContext.capabilities.drop` |  | `["ALL"]` |
| `controllerManager.kubeRbacProxy.image.repository` |  | `"gcr.io/kubebuilder/kube-rbac-proxy"` |
| `controllerManager.kubeRbacProxy.image.tag` |  | `"v0.14.1"` |
| `controllerManager.kubeRbacProxy.resources.limits.cpu` |  | `"500m"` |
| `controllerManager.kubeRbacProxy.resources.limits.memory` |  | `"128Mi"` |
| `controllerManager.kubeRbacProxy.resources.requests.cpu` |  | `"5m"` |
| `controllerManager.kubeRbacProxy.resources.requests.memory` |  | `"64Mi"` |
| `controllerManager.manager.args` |  | `["--health-probe-bind-address=:8081", "--metrics-bind-address=127.0.0.1:8080", "--leader-elect"]` |
| `controllerManager.manager.containerSecurityContext.allowPrivilegeEscalation` |  | `false` |
| `controllerManager.manager.containerSecurityContext.capabilities.drop` |  | `["ALL"]` |
| `controllerManager.manager.image.repository` |  | `"quay.io/validator-labs/validator-plugin-kubescape"` |
| `controllerManager.manager.image.tag` | x-release-please-version | `"v0.0.3"` |
| `controllerManager.manager.resources.limits.cpu` |  | `"500m"` |
| `controllerManager.manager.resources.limits.memory` |  | `"128Mi"` |
| `controllerManager.manager.resources.requests.cpu` |  | `"10m"` |
| `controllerManager.manager.resources.requests.memory` |  | `"64Mi"` |
| `controllerManager.manager.volumeMounts` |  | `[]` |
| `controllerManager.replicas` |  | `1` |
| `controllerManager.serviceAccount.annotations` |  | `{}` |
| `controllerManager.volumes` |  | `[]` |
| `kubernetesClusterDomain` |  | `"cluster.local"` |
| `metricsService.ports` |  | `[{"name": "https", "port": 8443, "protocol": "TCP", "targetPort": "https"}]` |
| `metricsService.type` |  | `"ClusterIP"` |
---
_Documentation generated by [Frigate](https://frigate.readthedocs.io)._

