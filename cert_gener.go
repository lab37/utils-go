package main
// 此程序用来生成证书，注意如果证书是由CA签发的，那么证书文件中
// 将同时包含服务器的签名以及CA的签名，其中服务器签名在前，CA签名在后
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

func main() {
	max := new(big.Int).Lsh(big.NewInt(1), 128) // 生成一个大数种子
	serialNumber, _ := rand.Int(rand.Reader, max) // 生成一个大随机数
	// 证书的组织名称，可以随便设置，只是一个显示名称
	subject := pkix.Name{
		Organization:       []string{"Manning Publications Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "Go Web Programming",
	}

	// 证书模板
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour), // 证书的有效期1年
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")}, //证书签证的IP，只能在这个IP上运行
	}

	pk, _ := rsa.GenerateKey(rand.Reader, 2048) // 生成2048位的RSA私钥和对应的公钥

	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk) // 生成一个经过DER编码格式化的字节切片
	certOut, _ := os.Create("cert.pem")  // 创建一个cert.pem文件
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})  // 将证书编码到了cert.pem中
	certOut.Close()

	keyOut, _ := os.Create("key.pem")  // 创建key.pem文件
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})  // 以PEM编码的方式把私钥编码并保存到key.pem中
	keyOut.Close()
}