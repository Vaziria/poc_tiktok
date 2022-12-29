package main

import (
	"github.com/vaziria/cremona/seller"
)

func main() {
	proxy := seller.NewConfigProxy()

	browser := seller.NewBrowser(proxy.Addr)
	defer browser.Service.Stop()
	go seller.StartProxy(proxy)

	service := seller.AkunService{
		ProxyListen: proxy.Addr,
		Browser:     browser,
	}
	service.GetAkunSession("ttest")
}
