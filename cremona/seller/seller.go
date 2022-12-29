package seller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SellerApi struct {
	Akun *Akun
}

type CustomerServiceInfo struct {
	PigeonCid  string `json:"pigeonCid"`
	ScreenName string `json:"screenName"`
	AvatarUrl  string `json:"avatarUrl"`
	Lang       string `json:"lang"`
	Status     int16  `json:"status"`
	OuterCid   string `json:"outerCid"`
}

type TccConfig struct {
	AllowUseVideo bool `json:"allowUseVideo"`
}

type ShopCsInfoData struct {
	ShopName            string              `json:"shopName"`
	PigeonShopId        string              `json:"pigeonShopId"`
	ShopLogo            string              `json:"shopLogo"`
	ShopRegion          string              `json:"shopRegion"`
	CustomerServiceInfo CustomerServiceInfo `json:"customerServiceInfo"`
	CustomerServiceType int16               `json:"customerServiceType"`
	RegionCode          string              `json:"regionCode"`
	OuterShopId         string              `json:"outerShopId"`
	ShopTimezone        string              `json:"shopTimezone"`
}

type Response[T any] struct {
	Code    int16  `json:"code"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

var defaultHeader = map[string][]string{
	"Accept":             {"application/json, text/plain, */*"},
	"Accept-Language":    {"en-US,en;q=0.9"},
	"Content-Type":       {"application/json"},
	"Referer":            {"https://seller-id.tiktok.com/chat"},
	"Sec-ch-ua":          {"\"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"108\", \"Google Chrome\";v=\"108\""},
	"Sec-ch-ua-mobile":   {"?0"},
	"Sec-ch-ua-platform": {"\"Windows\""},
	"Sec-Fetch-Dest":     {"empty"},
	"Sec-Fetch-Mode":     {"cors"},
	"Sec-Fetch-Site":     {"same-origin"},
	"User-Agent":         {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"},
}

func (s *SellerApi) GetShopAndCsInfo() ShopCsInfoData {

	url := "https://seller-id.tiktok.com/chat/api/seller/getShopAndCsInfo?PIGEON_BIZ_TYPE=1&aid=4068"
	var resdata Response[ShopCsInfoData]

	req, err := http.NewRequest("GET", url, nil)
	client := &http.Client{}

	if err != nil {
		panic("error create request")
	}

	for _, value := range s.Akun.Cookies {
		req.AddCookie(&value)
	}

	for key, value := range defaultHeader {
		req.Header.Set(key, value[0])
	}

	resp, err := client.Do(req)

	// error handle
	if err != nil {
		fmt.Printf("error = %s \n", err)
	}

	data, err := io.ReadAll(resp.Body)
	json.Unmarshal(data, &resdata)

	// error handle
	if err != nil {
		fmt.Printf("error = %s \n", err)
	}

	return resdata.Data

}

type TokenData struct {
	Token        string `json:"token"`
	Env          string `json:"env"`
	PigeonCid    string `json:"pigeonCid"`
	IdcRegion    string `json:"idcRegion"`
	WsUrl        string `json:"wsUrl"`
	BizServiceId int32  `json:"bizServiceId"`
	ApiUrl       string `json:"apiUrl"`
}

func (s *SellerApi) GetTokenInfo(sellerId string) TokenData {
	var resdata Response[TokenData]

	url := "https://seller-id.tiktok.com/chat/api/seller/token?PIGEON_BIZ_TYPE=1&oec_region=ID&aid=4068&oec_seller_id=" + sellerId
	req, err := http.NewRequest("GET", url, nil)
	client := &http.Client{}

	if err != nil {
		panic("error create request")
	}

	for _, value := range s.Akun.Cookies {
		req.AddCookie(&value)
	}

	for key, value := range defaultHeader {
		req.Header.Set(key, value[0])
	}

	resp, err := client.Do(req)

	// error handle
	if err != nil {
		fmt.Printf("error = %s \n", err)
	}

	data, err := io.ReadAll(resp.Body)
	json.Unmarshal(data, &resdata)

	// error handle
	if err != nil {
		fmt.Printf("error = %s \n", err)
	}

	return resdata.Data
}
