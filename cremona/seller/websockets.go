package seller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/gorilla/schema"
	"github.com/gorilla/websocket"
)

var encoder = schema.NewEncoder()
var appkey = "b42d99769353ce6304e74fb597e36e90"

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

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func NewSocketQuery() SocketQuery {
	return SocketQuery{
		WebsocketSwitchRegion: "ID",
		Aid:                   5341,
		Fpid:                  92,
		DevicePlatform:        "web",
		VersionCode:           10000,
	}
}

func (akun *Akun) createQuery() *SocketQuery {
	api := SellerApi{
		Akun: akun,
	}

	shopInfo := api.GetShopAndCsInfo()
	tokenInfo := api.GetTokenInfo(shopInfo.CustomerServiceInfo.OuterCid)
	query := NewSocketQuery()

	query.Token = tokenInfo.Token
	query.DeviceId = tokenInfo.PigeonCid

	contentHash := strconv.Itoa(int(query.Fpid)) + appkey + tokenInfo.PigeonCid + "f8a69f1719916z"
	query.AccessKey = GetMD5Hash(contentHash)

	return &query

}

func (akun *Akun) createUrl() *url.URL {
	query := akun.createQuery()
	q := url.Values{}

	err := encoder.Encode(query, q)

	if err != nil {
		fmt.Printf("error = %s \n", err)
	}

	u := url.URL{
		Scheme: "wss",
		Host:   "oec-im-frontier-va.tiktokglobalshop.com",
		Path:   "/ws/v2",
	}

	u.RawQuery = q.Encode()

	return &u
}

func (akun *Akun) CreateWebsocket() {
	u := akun.createUrl()

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	time.Sleep(10 * time.Second)

	defer c.Close()
}
