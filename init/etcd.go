package init

import (
	customConst "erbiaoOS/const"
	"log"
)

func EtcdCfg() {
	service := customConst.EtcdService
	log.Println(service)
}
