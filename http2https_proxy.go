package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"crypto/tls"
	"net/url"
)

type Pxy struct {}

func (p *Pxy) ServeHTTP(nRw http.ResponseWriter, oReq *http.Request) {
    fmt.Println("接收到的URL为:",oReq.URL)
	fmt.Println("接收到请求头中的Accept-Encoding为：",oReq.Header["Accept-Encoding"])
	fmt.Println("接收到请求头中的Content-Type为：",oReq.Header.Get("Content-Type"))
	fmt.Println("接收到请求的URI,Host,RemoteAddr分别为：",oReq.RequestURI, oReq.Host, oReq.RemoteAddr)
	fmt.Println("接收到请求体长度为:",oReq.ContentLength)
	oReq.ParseForm()
	fmt.Println(oReq.PostForm)
	fmt.Println(oReq.PostFormValue("firstname"))
	
	transport := &http.Transport{
        TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},  // go默认要检验服务器证书，此项配置为不让客户端检验证书
    }

	// 接收请求，并组建新请求
	targetUrl := "https://www.yiyaoshuju.cn"
    targetData := url.Values{"a":{"1"},"b":{"3"}}
    tr := &http.Transport{
        TLSClientConfig:    &tls.Config{InsecureSkipVerify: true}, //go默认要检验服务器证书，此项配置为不让客户端检验证书
    }
    client := &http.Client{Transport: tr}
	client.Jar, _ = cookiejar.New(nil)
    
	
	// 发出新请求，接收响应
	// resp, err := client.Get(targetUrl)
	resp, err := client.PostForm(targetUrl, targetData)
	if err != nil {
		nRw.WriteHeader(http.StatusBadGateway)
		return
	}
	// 处理响应，并把响应转发给客户端
	for key, value := range resp.Header {
		for _, v := range value {
			nRw.Header().Add(key, v)
		}
	}

	nRw.WriteHeader(resp.StatusCode)
	io.Copy(nRw, resp.Body)
	resp.Body.Close()
}

func main() {
	fmt.Println("代理端口 :8080")
	http.Handle("/", &Pxy{})
	http.ListenAndServe("127.0.0.1:8080", nil)
}