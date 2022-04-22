package cert

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"erbiaoOS/utils/file"
	"math/big"
	rd "math/rand"
	"net"
	"time"
)

// AltNames 可使用证书的域名、IP
type altNames struct {
	DNSNames []string
	IPs      []net.IP
}

// NewAltNames 初始化默认的域名、IP
func NewAltNames(IPs, dnsName []string) *altNames {
	dnsName = append(dnsName, "localhost", "localhost.localdomain", "localhost4", "localhost4.localdomain4")

	ips := []net.IP{net.ParseIP("127.0.0.1")}
	for _, ip := range IPs {
		if i := net.ParseIP(ip); i != nil {
			// ip添加到Subject Alternative Name - ip
			ips = append(ips, i)
		}
	}
	return &altNames{
		DNSNames: dnsName,
		IPs:      ips,
	}
}

// NewCertInfo 初始化证书信息
func newCertInfo(Oinfo []string, CNinfo string, IPs []net.IP, dnsNames []string) *x509.Certificate {
	var now = time.Now()

	return &x509.Certificate{
		SerialNumber: big.NewInt(rd.Int63()), //证书序列号
		Subject: pkix.Name{
			Country:            []string{"CN"},
			Organization:       Oinfo,
			OrganizationalUnit: []string{"system"},
			Province:           []string{"beijing"},
			CommonName:         CNinfo,
			Locality:           []string{"beijing"},
		},
		NotBefore:             now,                                                                        //证书有效期开始时间
		NotAfter:              now.AddDate(100, 0, 0),                                                     //证书有效期结束时间
		BasicConstraintsValid: true,                                                                       //基本的有效性约束
		IsCA:                  false,                                                                      //是否是根证书
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth}, //证书用途(客户端认证，数据加密)
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageDataEncipherment,
		IPAddresses:           IPs,
		DNSNames:              dnsNames,
	}
}

// Generate 生成证书信息
func generate(cer *x509.Certificate, caFile string) {
	// CA根证书可以通过openssl或cfssl工具生成，未使用go代码生成

	//解析CA根证书
	//caFile, err := ioutil.ReadFile("caPublicfile")
	//if err != nil {
	//	panic(err)
	//}
	//caBlock, _ := pem.Decode(caFile)

	caBlock, _ := pem.Decode([]byte(caPublicKey))

	cert, err := x509.ParseCertificate(caBlock.Bytes)
	if err != nil {
		panic(err)
	}
	//解析CA根私钥
	//keyFile, err := ioutil.ReadFile(caPrivatefile)
	//if err != nil {
	//	panic(err)
	//}
	//keyBlock, _ := pem.Decode(keyFile)

	keyBlock, _ := pem.Decode([]byte(caPrivateKey))
	praKey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		panic(err)
	}

	//生成公钥私钥对
	priKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	ca, err := x509.CreateCertificate(rand.Reader, cer, cert, &priKey.PublicKey, praKey)
	if err != nil {
		panic(err)
	}

	//编码证书文件和私钥文件
	caPem := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: ca,
	}
	ca = pem.EncodeToMemory(caPem)

	buf := x509.MarshalPKCS1PrivateKey(priKey)
	keyPem := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: buf,
	}

	key := pem.EncodeToMemory(keyPem)
	file.Create(caFile+".pem", string(ca))
	file.Create(caFile+"-key.pem", string(key))
}
