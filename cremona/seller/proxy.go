package seller

import (
	"log"
	"strings"

	"github.com/gorilla/schema"
	"github.com/kardianos/mitmproxy/cert"
	"github.com/kardianos/mitmproxy/proxy"
)

var SocketQChan = make(chan SocketQuery, 1)

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
func (addon *AuthGetter) Response(flow *proxy.Flow)                     {}
func (addon *AuthGetter) Request(flow *proxy.Flow) {
	url := flow.Request.URL.Path

	rawQuery := flow.Request.URL.Query()

	if strings.Contains(url, "/ws/v2") {
		log.Println("getting Query Websocket")

		var query SocketQuery

		decoder := schema.NewDecoder()

		err := decoder.Decode(&query, rawQuery)

		if err != nil {
			log.Println("Gagal Parse Query Websocket")
		} else {
			SocketQChan <- query
		}

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
