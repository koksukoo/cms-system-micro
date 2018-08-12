package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Proxy struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
}

func NewProxy(target string) *Proxy {
	url, _ := url.Parse(target)
	return &Proxy{target: url, proxy: httputil.NewSingleHostReverseProxy(url)}
}

func (p *Proxy) handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Forwarded-Host", r.Header.Get("Host"))
	// r.URL.Host = p.target.Host
	// r.URL.Scheme = p.target.Scheme
	// r.Host = p.target.Host
	p.proxy.ServeHTTP(w, r)
}
