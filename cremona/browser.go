package main

import (
	"path/filepath"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type Browser struct {
	ProxyAddr string
	Service   *selenium.Service
}

func (b *Browser) CreateDriver(profile string) selenium.WebDriver {

	proxy := "--proxy-server=" + b.ProxyAddr
	pathAbs, _ := filepath.Abs(profile)
	profileArg := "--user-data-dir=" + pathAbs

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{
		Args: []string{
			proxy,
			profileArg,
		},
	})

	driver, err := selenium.NewRemote(caps, "")

	if err != nil {
		panic(err)
	}

	return driver
}

func NewBrowser(proxy string) *Browser {
	service, err := selenium.NewChromeDriverService("./chromedriver", 4444)

	if err != nil {
		panic(err)
	}

	browser := &Browser{
		ProxyAddr: proxy,
		Service:   service,
	}

	return browser
}
