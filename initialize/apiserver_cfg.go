package initialize

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/pkg/file"
	"strings"
)

func KubeApiserverCfg(masterHost []string, etcdServerUrls string) {
	for _, ip := range masterHost {
		cfg := strings.ReplaceAll(customConst.KubeApiserverCfg, "currentIPaddr", ip)
		cfg = strings.ReplaceAll(cfg, "etcdServerUrls", etcdServerUrls)
		file.Create(customConst.K8sMasterCfgDir+ip+"-kube-apiserver", cfg)
	}
}
