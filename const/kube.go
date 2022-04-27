package myConst

const (
	SetClusterCmd = "kubectl config set-cluster kubernetes --certificate-authority=%s --embed-certs=true --server=https://%s:6443 --kubeconfig=%s"

	SetCredentialsCmd = "kubectl config set-credentials %s --client-certificate=%s --client-key=%s --embed-certs=true --kubeconfig=%s"

	KubeletSetCredentialsCmd = "kubectl config set-credentials %s --token=%s --kubeconfig=%s"

	UseContextCmd = "kubectl config use-context %s --kubeconfig=%s"

	SetContextCmd            = "kubectl config set-context %s --cluster=kubernetes --user=%s --kubeconfig=%s"
	ClusterrolebindingDelete = "kubectl delete clusterrolebinding %s || echo 不存在此clusterrolebinding"
	ClusterrolebindingCreate = "kubectl create clusterrolebinding %s --clusterrole=%s --user=%s"
)
