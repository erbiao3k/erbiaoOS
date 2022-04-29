package kubectl

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/pkg/cert"
	"erbiaoOS/pkg/config"
	"fmt"
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

var (

	// kubectl 管理客户端初始化指令
	setClusterCmd            = fmt.Sprintf(myConst.SetClusterCmd, cert.CaPubilcKeyFile, config.ApiserverEnterpoint, kubeconfig)
	setCredentialsCmd        = fmt.Sprintf(myConst.SetCredentialsCmd, user, publicKeyFile, privateKeyFile, kubeconfig)
	setContextCmd            = fmt.Sprintf(myConst.SetContextCmd, context, user, kubeconfig)
	useContextCmd            = fmt.Sprintf(myConst.UseContextCmd, context, kubeconfig)
	clusterrolebindingDelete = fmt.Sprintf(myConst.ClusterrolebindingDelete, credentials)
	clusterrolebindingCreate = fmt.Sprintf(myConst.ClusterrolebindingCreate, clusterrolebinding, clusterrole, clusterrolebindingUser)
)
