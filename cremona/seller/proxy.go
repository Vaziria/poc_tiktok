package seller

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/kardianos/mitmproxy/cert"
	"github.com/kardianos/mitmproxy/proxy"
)

var TokenChan = make(chan TokenData, 1)

type ConfigProxy struct {
	Addr string
}

type AuthGetter struct {
	proxy.BaseAddon
}

func (addon *AuthGetter) ClientConnected(client *proxy.ClientConn)      {}
func (addon *AuthGetter) ClientDisconnected(client *proxy.ClientConn)   {}
func (addon *AuthGetter) ServerConnected(connCtx *proxy.ConnContext)    {}
func (addon *AuthGetter) ServerDisconnected(connCtx *proxy.ConnContext) {}
func (addon *AuthGetter) Requestheaders(f *proxy.Flow)                  {}
func (addon *AuthGetter) Request(flow *proxy.Flow)                      {}
func (addon *AuthGetter) Response(flow *proxy.Flow) {
	u := flow.Request.URL.Path

	if strings.Contains(u, "/seller/token") {
		var resdata Response[TokenData]

		json.Unmarshal(flow.Response.Body, &resdata)
		TokenChan <- resdata.Data

	}

}

func NewConfigProxy() *ConfigProxy {
	config := ConfigProxy{
		Addr: "127.0.0.1:6002",
	}

	return &config
}

func StartProxy(config *ConfigProxy) {
	var authGetter proxy.Addon = &AuthGetter{}

	certloader, err := cert.NewPathLoader("")
	if err != nil {
		log.Fatal(err)
	}

	ca, err := cert.New(certloader)
	if err != nil {
		log.Fatal(err)
	}

	opts := &proxy.Options{
		Addr:                  config.Addr,
		StreamLargeBodies:     1024 * 1024 * 5,
		InsecureSkipVerifyTLS: false,
		CA:                    ca,
	}

	p, err := proxy.NewProxy(opts)
	if err != nil {
		log.Fatal(err)
	}

	// p.AddAddon(&addon.LogAddon{})
	p.AddAddon(authGetter)

	err = p.Start()

	if err != nil {
		panic("error start proxy")
	}
}
