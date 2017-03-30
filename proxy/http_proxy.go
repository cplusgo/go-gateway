package proxy

import (
	"net/http"
	"net/url"
	"log"
	"io"
	"io/ioutil"
)

/**
 * http请求代理，转发到具体的业务服务器
 * 注意这里把业务服务器当做代理服务器了，
 * 但是业务服务器并不转发该请求，而是直接处理了请求
 */
type HttpProxy struct {
	transport *http.Transport
}

func NewHttpProxy(rawurl string) *HttpProxy {
	_, err := url.Parse(rawurl)
	if err != nil {
		log.Println(rawurl, "NewHttpProxy failed!")
		return nil
	}
	transport := &http.Transport{
		//Proxy: http.ProxyURL(url),
	}
	proxy := &HttpProxy{transport: transport}
	return proxy
}

func (this *HttpProxy) ServeRequest(w http.ResponseWriter, req *http.Request, rawurl string) {
	client := &http.Client{
		Transport: this.transport,
	}
	for _, cookie := range req.Cookies() {
		req.AddCookie(cookie)
	}
	req.RequestURI = ""
	u, _ := url.Parse(rawurl)
	req.URL = u
	req.Host = u.Host
	for k, _ := range req.Header {
		if k != "Cookie" {
			req.Header.Del(k)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return
	}
	io.Copy(w, resp.Body)
	resp.Body.Close()
}

func (this *HttpProxy) ProxyRequest(req *http.Request) []byte {
	client := &http.Client{
		Transport: this.transport,
	}
	resp, _ := client.Do(req)
	data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return data

}
