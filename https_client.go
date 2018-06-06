package main
// 更多的内容参见https://studygolang.com/articles/2946文章名《Go和HTTPS〉
import (
    "crypto/tls"
    "fmt"
    "io/ioutil"
    "net/http"
	"net/http/cookiejar"
    // "net/url"

)

func main() {

    targetUrl := "https://www.yiyaoshuju.cn"
    // targetData := url.Values{"a":{"1"},"b":{"3"}}
    tr := &http.Transport{
        TLSClientConfig:    &tls.Config{InsecureSkipVerify: true}, //go默认要检验服务器证书，此项配置为不让客户端检验证书
    }
    client := &http.Client{Transport: tr}
	client.Jar, _ = cookiejar.New(nil)
    resp, err := client.Get(targetUrl)
	// resp, err := client.PostForm(targetUrl, targetData)


    if err != nil {
        fmt.Println("error:", err)
        return
    }
	fmt.Println(resp)
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}