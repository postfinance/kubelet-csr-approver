# kubelet-csr-approver

Kubelet CSR approver is a Kubernetes controller whose sole purpose is to
auto-approve [`kubelet-serving` Certificate Signing Request(CSR)](https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/kubeadm-certs/#kubelet-serving-certs),
provided these CSRs comply with a series of configurable, provider-specific,
checks/verifications.

ℹ more details (installation, etc.) can be found [on the main
repository](https://github.com/postfinance/kubelet-csr-approver)
