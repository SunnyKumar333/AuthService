package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ProxyToService(baseURL string, pathPrefix string) http.HandlerFunc {
	target, err := url.Parse(baseURL)
	if err != nil {
		// utils.WriteJSONErrorResponse()
		fmt.Println(err.Error())
		return nil
	}
	// fmt.Println("Target:", target.Host, target.Path)
	proxy := httputil.NewSingleHostReverseProxy(target)

	originalDiractor := proxy.Director

	proxy.Director = func(r *http.Request) {
		originalDiractor(r)
		// fmt.Println(r.Host, r.URL.Host, r.URL.Path)
		r.Host = target.Host
		r.URL.Path = strings.TrimPrefix(r.URL.Path, pathPrefix)

		if userId, ok := r.Context().Value("userId").(string); ok {
			r.Header.Set("X-userId", userId)
		}

	}

	return proxy.ServeHTTP
}
