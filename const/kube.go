package customConst

const (
	SetClusterCmd = "kubectl config set-cluster kubernetes --certificate-authority=%s --embed-certs=true --server=https://%s:6443 --kubeconfig=%s\n"

	SetCredentialsCmd = "kubectl config set-credentials %s --client-certificate=%s --client-key=%s --embed-certs=true --kubeconfig=%s\n"

	UseContextCmd = "kubectl config use-context %s --kubeconfig=%s\n"

	SetContextCmd = "kubectl config set-context %s --cluster=kubernetes --user=%s --kubeconfig=%s\n"
)
