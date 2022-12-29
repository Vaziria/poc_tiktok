package seller

import (
	"net/http"
	"time"

	"github.com/tebeka/selenium"
)

var locSession = "data/sessions/"
var locProfile = "data/profiles/"

type Akun struct {
	Cookies []http.Cookie
}

type AkunService struct {
	ProxyListen string
	Browser     *Browser
}

func (s *AkunService) getAkunFromSelenium(name string) *Akun {

	profile := locProfile + name
	driver := s.Browser.CreateDriver(profile)
	defer driver.Close()

	driver.Get("https://seller-id.tiktok.com/chat")
	time.Sleep(5 * time.Second)

	cookies := getHttpCookies(driver)

	akun := &Akun{
		Cookies: cookies,
	}

	// <-SocketQChan

	return akun
}

func (s *AkunService) GetAkunSession(name string) *Akun {
	return s.getAkunFromSelenium(name)
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
