package main

func main() {
	proxy := NewConfigProxy()

	browser := NewBrowser(proxy.Addr)
	defer browser.Service.Stop()
	go StartProxy(proxy)

	service := AkunService{
		ProxyListen: proxy.Addr,
		Browser:     browser,
	}

	service.GetAkunSession("ttest")
}
