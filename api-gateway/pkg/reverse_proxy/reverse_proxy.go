package reverse_proxy

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ReverseProxy(ctx *gin.Context) {
	path := ctx.Request.URL.Path
	target, err := Target(path)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	if targetUrl, err := url.Parse(target); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	} else {
		Proxy(targetUrl).ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func Target(path string) (string, error) {
	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
	if len(parts) <= 1 {
		return "", fmt.Errorf("failed to parse target host from path: %s", path)
	}
	targetHost := fmt.Sprintf("svc-%s", parts[1])
	targetNamespace := fmt.Sprintf("svc-%s", parts[2])
	if targetHost == "" {
		return "", fmt.Errorf("failed to parse target host from path: %s", path)
	}
	targetAddr := fmt.Sprintf(
		"http://%s.%s:%d/%s",
		targetHost, targetNamespace, 10000, strings.Join(parts[3:], "/"),
	)
	return targetAddr, nil
}

func Proxy(address *url.URL) *httputil.ReverseProxy {
	p := httputil.NewSingleHostReverseProxy(address)
	p.Director = func(request *http.Request) {
		request.Host = address.Host
		request.URL.Scheme = address.Scheme
		request.URL.Host = address.Host
		request.URL.Path = address.Path
	}
	p.ModifyResponse = func(response *http.Response) error {
		if response.StatusCode == http.StatusInternalServerError {
			u, s := readBody(response)
			log.Fatalf("%s ,req %s ,with error %d, body:%s", u.String(), address, response.StatusCode, s)
			response.Body = ioutil.NopCloser(bytes.NewReader([]byte(fmt.Sprintf("error %s", u.String()))))
		} else if response.StatusCode > 300 {
			_, s := readBody(response)
			log.Fatalf("req %s ,with error %d, body:%s", address, response.StatusCode, s)
			response.Body = ioutil.NopCloser(bytes.NewReader([]byte(s)))
		}
		return nil
	}
	return p
}

func readBody(response *http.Response) (uuid.UUID, string) {
	defer response.Body.Close()
	all, _ := ioutil.ReadAll(response.Body)
	u := uuid.New()
	var s string
	if len(all) > 0 {
		s = string(all)
	}
	return u, s
}
