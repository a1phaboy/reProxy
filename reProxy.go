package main

import (
	"github.com/pretty66/websocketproxy"
	"net/http"
)

func main() {
	PHPwp, err := websocketproxy.NewProxy("ws://127.0.0.1:8009", func(r *http.Request) error {
		// 权限验证
		r.Header.Set("Cookie", "----")
		// 伪装来源
		r.Header.Set("Origin", "http://127.0.0.1:8009")
		return nil
	})
	JAVAwp, err := websocketproxy.NewProxy("ws://127.0.0.1:8080/x", func(r *http.Request) error {
		// 权限验证
		r.Header.Set("Cookie", "----")
		// 伪装来源
		r.Header.Set("Origin", "http://127.0.0.1:8080")
		return nil
	})
	if err != nil {

	}
	// 代理路径
	http.HandleFunc("/php_wsproxy", PHPwp.Proxy)
	http.HandleFunc("/java_wsproxy", JAVAwp.Proxy)
	http.ListenAndServe(":9696", nil)
}
