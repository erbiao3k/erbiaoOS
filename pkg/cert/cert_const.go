package cert

import (
	myConst "erbiaoOS/const"
)

const (
	CaPrivateKeyFile = myConst.CaCenterDir + "ca-key.pem"
	CaPubilcKeyFile  = myConst.CaCenterDir + "ca.pem"

	// caPublicKey ca机构证书公钥信息
	caPublicKey = `-----BEGIN CERTIFICATE-----
MIIDnDCCAoSgAwIBAgIUBgmIRptWdX9nH0+AALFJ2bw11AMwDQYJKoZIhvcNAQEL
BQAwZTELMAkGA1UEBhMCQ04xEDAOBgNVBAgTB0JlaWppbmcxEDAOBgNVBAcTB0Jl
aWppbmcxDDAKBgNVBAoTA2s4czEPMA0GA1UECxMGc3lzdGVtMRMwEQYDVQQDEwpr
dWJlcm5ldGVzMCAXDTIyMDQyMTAyNTAwMFoYDzIyMjIwMzA0MDI1MDAwWjBlMQsw
CQYDVQQGEwJDTjEQMA4GA1UECBMHQmVpamluZzEQMA4GA1UEBxMHQmVpamluZzEM
MAoGA1UEChMDazhzMQ8wDQYDVQQLEwZzeXN0ZW0xEzARBgNVBAMTCmt1YmVybmV0
ZXMwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDf2B+d8aGNoRYZI1zV
Hinan/2CLQAhu/h6D2HToRivvCftYMs0zy9kz4DREOivBUEAHsUIQhma8KWM0BBx
1U0CWB1FY+CVGFaZNEooTWnbA1RMLBssbOL6+HghgD/KacGd2QT91l6geJyMzrUi
j34bWcg8Ssbd9NwlbTk6/0qeGCzi+B3aT6pgVJG9rcJEG+e48Jj7CnedB4/zYSba
VBozlTQ4nchBU5dsgP/eLYpKq6wS9Hx2VYD+1Xe6sQ1PHuBY+mR02NdCKnK56Dlq
7prY+2bXZbVwCo4ThxCgQv0Ne5jffOLWTCPgMgFKmBhxXIrN2tqAy/2t2jC5WsbE
BeExAgMBAAGjQjBAMA4GA1UdDwEB/wQEAwIBBjAPBgNVHRMBAf8EBTADAQH/MB0G
A1UdDgQWBBT6zxNpJplUM36xZfhd8wgKGAZpzzANBgkqhkiG9w0BAQsFAAOCAQEA
PVPGo30FwYpiVnFHNVYuaN0sQ2q1kQwZrVnuRklmNXJpiEKM/fByAgenoknXTHXD
WfS2g8KzYi2JOLUHYvGjew/oLET3JcQJ8kXeGl7ejZHaIRmazK2NbUia553knpaO
SRUayXBpMr2/ndEQWFg91M2eNS13ncngTaI653qnWpVgktLjQkWyQtaSPeK51HNL
iRvm+XWTOysAXbjVj+MPzDIbcRWHQ91RnJKqIU/yzFdjHbO+IBuyf8vhjImRdVsv
iBaGhqh18oxcmdqUMf4nlVFYNzhisciVwxe6GarJ9SuPK5jTOK0sdkU7QEJd5AAv
PPdEyFDeV8uu1ebjVqDe8w==
-----END CERTIFICATE-----
`
	// caPrivateKey ca机构证书私钥信息
	caPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA39gfnfGhjaEWGSNc1R4p2p/9gi0AIbv4eg9h06EYr7wn7WDL
NM8vZM+A0RDorwVBAB7FCEIZmvCljNAQcdVNAlgdRWPglRhWmTRKKE1p2wNUTCwb
LGzi+vh4IYA/ymnBndkE/dZeoHicjM61Io9+G1nIPErG3fTcJW05Ov9Knhgs4vgd
2k+qYFSRva3CRBvnuPCY+wp3nQeP82Em2lQaM5U0OJ3IQVOXbID/3i2KSqusEvR8
dlWA/tV3urENTx7gWPpkdNjXQipyueg5au6a2Ptm12W1cAqOE4cQoEL9DXuY33zi
1kwj4DIBSpgYcVyKzdragMv9rdowuVrGxAXhMQIDAQABAoIBAQDd87fN7aSqAXse
6/CFpUYM8Kz24dYKnQ7RQOVYaZlHz0Kr3lk/gNxWkmYBZ0nBGW2NR/VIrfojptAQ
YtKbfMvCMkq73j/2tk5P1QhfE/uNXay9ZtJ+52zdO3gqh7c45kpEUMbhRylG5rEb
8W6r2SpKxmiEWAT+WmfbeITR2gdL34MPeNyX7aIhSE+sCyaDSPhL2pn7AAdVTnPZ
Pk2wPu34qOVi1X6/Enwzyh97qCeYmmx/CIAuptbAp/428V4x0lkEF6vWNcZIxv5f
vN+z3Bl2rvZcw/5oTauj6qHvSWqEQKaTW1My+oWbhnxi4lYiH6lQQXjAuM8nhpep
i9QRPrQ9AoGBAOBOVWU2qKeeb3ZOn6bx1K+TZvK4omTCegGVE2SNOnzEOvW919ZX
0By4hhLA7wMgqNt076PEzc4iacSnts7yqOlyrL3IAVKSPdSApxBHDCgh+Zv9x1n1
+cNpdlS+IENyogv/j5gQ+5u2Q+kfHpH4fhDY9KLt/43PN4YhUOppX3wnAoGBAP95
Fko4VvjNea5sOCiaiNJGqEugBUkfcpf4WV5y2oz9PN77wbAi3B/ql5I4HfCD08/8
MkmjWw8HBPlMvW/fOOMgJerqXtF2N686QRKCQ4UU5jAsxJ3zXLvETCBUMcJG6hG6
6GyzTefnsu50/6Ah1q6u4S2uPCRLCxj+YXHvspbnAoGAePimQ5Tr9qKS+JpErlPE
YgC8R/Fd27uq80mEEPm97mYiakA9tKLdYW//FwQoo2Ysy1bQm2FboW2b32yYQhpL
EMRA94VzeSXX5NCRyUyX+NkB9qgyqIjpcANjxyZW3ilnzdLBjcCzAfKVw5d99Dmx
O8LWhTyYU9HK3zL+ob28uocCgYAaFCBEX0/xgfgj6AQrkOranD/dyG4BsuYdwUpO
K+dHcSpfkM+KzWQvFeF3Gadkv/BFUPdJMRXAiPTnBgBohR7ngIaeXmJje2/fwVCX
NRjzYtjEni1L+mXC/RzQSAf0Twzh1nSXdA5F2A8Z7HOTwyCJIGz4Hssg4VA2svD7
kn5mjQKBgGYy/ffJo1KcP5iKn5RuzuTbJt8GDvZk9M/aSG5/XdGnV4AgU6gjSuZv
tjCdO3eNwMO9ayt9okvhaaobSKWuLhxNbnIgocDg/sYNrHSBVLZ94bijFZe3lTIo
Ms47mnxmnYez1ZlczmZqb+dess34EDnz/3M33k6g5Hqy2OR59656
-----END RSA PRIVATE KEY-----
`
)

/*
CA机构证书的生成规则，有效期200年
cat > ca-config.json << EOF
{
    "signing": {
        "default": {
            "expiry": "1752000h"
        },
        "profiles": {
            "kubernetes": {
                "expiry": "1752000h",
                "usages": [
                    "signing",
                    "key encipherment",
                    "server auth",
                    "client auth"
                ]
            }
        }
    }
}
EOF

cat > ca-csr.json << EOF
{
  "CN": "kubernetes",
  "key": {
      "algo": "rsa",
      "size": 2048
  },
  "names": [
    {
      "C": "CN",
      "ST": "Beijing",
      "L": "Beijing",
      "O": "k8s",
      "OU": "system"
    }
  ],
  "ca": {
          "expiry": "1752000h"
  }
}
EOF

根证书生成指令
cfssl gencert -initca ca-csr.json | cfssljson -bare ca -
*/
