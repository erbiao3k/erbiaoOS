package init

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/pkg/file"
)

// CfsslBinary cfssl二进制文件路径
func CfsslBinary() string {
	return customConst.TempData + file.ListHasPrefix(customConst.TempData, []string{"cfssl_1.6.1"})[0]
}

// CfssljsonBinary cfssljsonBinary二进制文件路径
func CfssljsonBinary() string {
	return customConst.TempData + file.ListHasPrefix(customConst.TempData, []string{"cfssljson_1.6.1"})[0]
}

// CfsslcertinfoBinary CfsslcertinfoBinary二进制文件路径
func CfsslcertinfoBinary() string {
	return customConst.TempData + file.ListHasPrefix(customConst.TempData, []string{"cfssl-certinfo_1.6.1"})[0]
}
