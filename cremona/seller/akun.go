package seller

import (
	"net/http"

	"github.com/tebeka/selenium"
)

type SocketQuery struct {
	Token                 string `schema:"token"`
	Aid                   int32  `schema:"aid"`
	Fpid                  int32  `schema:"fpid"`
	DeviceId              string `schema:"device_id"`
	AccessKey             string `schema:"access_key"`
	DevicePlatform        string `schema:"device_platform"`
	VersionCode           int32  `schema:"version_code"`
	WebsocketSwitchRegion string `schema:"websocket_switch_region"`
}

type Akun struct {
	Token   string
	Cookies []http.Cookie
	QueryWs SocketQuery
}

type AkunService struct {
	ProxyListen string
	Browser     *Browser
}

func (s *AkunService) GetAkunSession(profile string) *Akun {

	driver := s.Browser.CreateDriver(profile)
	defer driver.Close()

	driver.Get("https://seller-id.tiktok.com/chat")

	cookies := getHttpCookies(driver)

	akun := &Akun{
		Cookies: cookies,
		// QueryWs: socketQuery,
	}

	service := SellerService{
		Akun: akun,
	}

	service.GetShopAndCsInfo()
	service.GetTokenInfo()

	<-SocketQChan
	return akun
}

func CreateAkunService() *AkunService {

	service := AkunService{}

	return &service
}

func getHttpCookies(driver selenium.WebDriver) []http.Cookie {

	seleCookies, err := driver.GetCookies()

	if err != nil {
		panic("error getting cookies")
	}

	httpCookies := make([]http.Cookie, len(seleCookies))

	for key, cookie := range seleCookies {

		httpCookies[key] = http.Cookie{
			Name:  cookie.Name,
			Value: cookie.Value,
			Path:  cookie.Path,
		}
	}

	return httpCookies

}
