package proxy

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
	"net/url"
	"os"
)

var proxyUrl = os.Getenv("HTTP_PROXY_URL")
var proxyUser = os.Getenv("HTTP_PROXY_USER")
var proxyPass = os.Getenv("HTTP_PROXY_PASS")

func GetProxyUrl() func(*http.Request) (*url.URL, error) {
	if proxyURL, err := url.Parse(proxyUrl); err != nil {
		log.Fatalln(err)
	} else {
		if proxyUser != "" && proxyPass != "" {
			proxyURL.User = url.UserPassword(proxyUser, proxyPass)
		}
		return http.ProxyURL(proxyURL)
	}

	return nil
}
