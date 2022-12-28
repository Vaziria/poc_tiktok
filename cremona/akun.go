package main

import (
	"net/http"

	"github.com/tebeka/selenium"
)

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

	socketQuery := <-SocketQChan

	akun := &Akun{
		Cookies: cookies,
		QueryWs: socketQuery,
	}

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
