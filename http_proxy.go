package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

type Pxy struct {}

func (p *Pxy) ServeHTTP(nRw http.ResponseWriter, oReq *http.Request) {
	fmt.Printf("Received request %s %s %s\n", oReq.Method, oReq.Host, oReq.RemoteAddr)

	transport :=  http.DefaultTransport

	// 接收请求，并组建新请求
	nReq := new(http.Request)
	*nReq = *oReq // 不对请求做任何改变，直接复制到新请求中

	if clientIP, _, err := net.SplitHostPort(oReq.RemoteAddr); err == nil {
		if prior, ok := nReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		nReq.Header.Set("X-Forwarded-For", clientIP)
	}

	// 发出新请求，接收响应
	response, err := transport.RoundTrip(nReq)
	if err != nil {
		nRw.WriteHeader(http.StatusBadGateway)
		return
	}

	// 处理响应，并把响应转发给客户端
	for key, value := range response.Header {
		for _, v := range value {
			nRw.Header().Add(key, v)
		}
	}

	nRw.WriteHeader(response.StatusCode)
	io.Copy(nRw, response.Body)
	response.Body.Close()
}

func main() {
	fmt.Println("代理端口 :8080")
	http.Handle("/", &Pxy{})
	http.ListenAndServe("127.0.0.1:8080", nil)
}