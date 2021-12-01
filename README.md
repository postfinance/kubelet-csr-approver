# kubelet-csr-approver

Kubelet CSR approver is a Kubernetes controller whose sole purpose is to auto-approve [`kubelet-serving`
Certificate Signing Request (CSR)](https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/kubeadm-certs/#kubelet-serving-certs),
provided these CSRs comply with a series of configurable, provider-specific, checks/verifications.  

ℹ more details can be found [on the main repository](https://github.com/postfinance/kubelet-csr-approver)

## Helm Install

Adjust `providerRegex` as needed.

```sh
helm repo add kubelet-csr-approver https://postfinance.github.io/kubelet-csr-approver
helm install kubelet-csr-approver kubelet-csr-approver/kubelet-csr-approver -n kube-system \
  --set providerRegex='^node-\w*\.int\.company\.ch$'
```

