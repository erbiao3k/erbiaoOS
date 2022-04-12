package cert

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

// GenerateCert 生成证书
func GenerateCert(host []string, commonName, certDir, CertName string) {
	//1.生成密钥对
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	//2.创建证书模板
	pkiName := pkix.Name{
		Country:            []string{"CN"},
		Organization:       []string{"k8s"},
		OrganizationalUnit: []string{"system"},
		Locality:           []string{"beijing"},
		Province:           []string{"beijing"},
		CommonName:         commonName,
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1), //该号码表示CA颁发的唯一序列号，在此使用一个数来代表

		Issuer:      pkiName,
		Subject:     pkiName,
		NotBefore:   time.Now(),
		NotAfter:    time.Now().AddDate(100, 0, 0),
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign, //表示该证书是用来做服务端认证的
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}

	for _, h := range host {
		if ip := net.ParseIP(h); ip != nil {
			// ip添加到Subject Alternative Name - ip
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			// 域名添加到Subject Alternative Name - 域名，主机名
			template.DNSNames = append(template.DNSNames, h)
		}
	}

	template.IPAddresses = append(template.IPAddresses, net.ParseIP("127.0.0.1"))
	template.DNSNames = append(template.DNSNames, "localhost", "localhost.localdomain", "localhost4", "localhost6", "localhost4.localdomain4")

	//3.创建证书,这里第二个参数和第三个参数相同则表示该证书为自签证书，返回值为DER编码的证书
	certificate, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		panic(err)
	}

	//4.将得到的证书放入pem.Block结构体中
	block := pem.Block{
		Type:    "CERTIFICATE",
		Headers: nil,
		Bytes:   certificate,
	}

	//5.通过pem编码并写入磁盘文件
	file, err := os.Create(certDir + "/" + CertName + ".pem")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	pem.Encode(file, &block)

	//6.将私钥中的密钥对放入pem.Block结构体中
	block = pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   x509.MarshalPKCS1PrivateKey(priv),
	}

	//7.通过pem编码并写入磁盘文件
	file, err = os.Create(certDir + "/" + CertName + "-key.pem")
	if err != nil {
		panic(err)
	}
	pem.Encode(file, &block)
}
