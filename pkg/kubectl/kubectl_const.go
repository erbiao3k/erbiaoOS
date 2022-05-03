package kubectl

import (
	myConst "erbiaoOS/vars"
)

const (
	user                   = "admin"
	context                = "kubernetes"
	credentials            = context
	clusterrolebinding     = context
	clusterrolebindingUser = context

	clusterrole    = "cluster-admin"
	kubeconfig     = myConst.KubectlConfigDir + "config"
	privateKeyFile = myConst.K8sSslDir + "admin-key.pem"
	publicKeyFile  = myConst.K8sSslDir + "admin.pem"
)
