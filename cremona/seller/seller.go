package seller

import (
	"fmt"
	"io"
	"net/http"
)

type SellerService struct {
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

type ShopCsInfoRes struct {
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

func (s *SellerService) GetShopAndCsInfo() {

	url := "https://seller-id.tiktok.com/chat/api/seller/getShopAndCsInfo?PIGEON_BIZ_TYPE=1&aid=4068"
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

	// error handle
	if err != nil {
		fmt.Printf("error = %s \n", err)
	}

	// Print response
	fmt.Printf("Response = %s", string(data))

}

// {
//     "token": "DlZTLo19XB0MxvOHyWIWtHRKUz41mGlzXbUhNvDcIWWvGkG6FUdx43",
//     "env": "",
//     "pigeonCid": "7080217658917961989",
//     "idcRegion": "maliva",
//     "wsUrl": "wss://oec-im-frontier-va.tiktokglobalshop.com/ws/v2",
//     "bizServiceId": 20041,
//     "apiUrl": "https://imapi-va-oth.isnssdk.com/"
// }

func (s *SellerService) GetTokenInfo() {
	url := "https://seller-id.tiktok.com/chat/api/seller/token?PIGEON_BIZ_TYPE=1&oec_region=ID&aid=4068&oec_seller_id=7494567309891832821"
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

	// error handle
	if err != nil {
		fmt.Printf("error = %s \n", err)
	}

	// Print response
	fmt.Printf("Response = %s", string(data))
}
