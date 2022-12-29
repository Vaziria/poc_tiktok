package seller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gorilla/schema"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

var encoder = schema.NewEncoder()
var appkey = "b42d99769353ce6304e74fb597e36e90"

func cookieString(cookies []http.Cookie) string {
	var hasil string

	for _, value := range cookies {
		hasil += value.Name + "=" + value.Value + ";"
	}

	return hasil

}

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

	query := NewSocketQuery()

	query.Token = akun.Token
	query.DeviceId = akun.DeviceId

	contentHash := strconv.Itoa(int(query.Fpid)) + appkey + akun.DeviceId + "f8a69f1719916z"
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

func (akun *Akun) createPing() []byte {

	// data := "CJJOELPNmOnVMBiSTiABOgJwYkKGAQjIARCSThoFMC4zLjgiNkpLc2k3M3Z3aFdObjllQkNRSW9WY3l0OUhhZm5SMVJlT1dvYjg1NGtGSHNRSTBzcEJ1MnhLdCgDMAA6DjEyYzkyOWE6bWFzdGVyQg7CDAsIquPM+t2q+AIQMkoTNzA4MDIxNzY1ODkxNzk2MTk4OVoDd2VikAEC"

	// hasil, _ := base64.StdEncoding.DecodeString(data)
	// 	headers
	// :
	// []
	// logid
	// :
	// n {low: 1565413085, high: 389, unsigned: false}
	// method
	// :
	// 1
	// payload
	// :
	// Uint8Array(134) [8, 200, 1, 16, 146, 78, 26, 5, 48, 46, 51, 46, 56, 34, 54, 70, 119, 65, 55, 89, 90, 68, 82, 73, 51, 89, 76, 113, 114, 107, 76, 57, 73, 73, 72, 105, 65, 105, 112, 109, 81, 57, 112, 107, 106, 115, 70, 106, 75, 113, 67, 88, 97, 110, 52, 86, 56, 71, 88, 84, 115, 119, 120, 101, 109, 99, 100, 119, 75, 40, 3, 48, 0, 58, 14, 49, 50, 99, 57, 50, 57, 97, 58, 109, 97, 115, 116, 101, 114, 66, 14, 194, 12, 11, 8, 225, 210, 175, 242, 208, …]
	// payload_type
	// :
	// "pb"
	// seqid
	// :
	// n {low: 10002, high: 0, unsigned: false}
	// service
	// :
	// 10002

	data := &Request{
		Headers:        map[string]string{},
		AuthType:       2,
		DevicePlatform: "web",
		InboxType:      0,
		BuildNumber:    "12c929a:master",
		SdkVersion:     "0.3.8",
		Cmd:            203,
		Body: &RequestBody{
			MessagesPerUserInitV2Body: &MessagesPerUserInitV2RequestBody{
				Cursor: 0,
			},
		},
		Refer:      3,
		Token:      akun.Token,
		DeviceId:   akun.DeviceId,
		SequenceId: 10001,
	}

	payload, err := proto.Marshal(data)

	frame := &Frame{
		Seqid:       10001,
		Logid:       10001,
		Service:     10002,
		PayloadType: "pb",
		Method:      1,
		Payload:     payload,
	}

	if err != nil {
		panic("gagal create request websocket")
	}

	hasil, _ := proto.Marshal(frame)

	log.Println("creating ping", fmt.Sprintf("%x", hasil))

	return hasil
}

func (akun *Akun) CreateWebsocket() {
	u := akun.createUrl()

	requester := http.Header{
		"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"},
		"Origin":     {"https://seller-id.tiktok.com"},
		"Cookie":     {cookieString(akun.Cookies)},
	}

	c, res, err := websocket.DefaultDialer.Dial(u.String(), requester)

	log.Println(res.StatusCode)

	if err != nil {
		log.Fatal("dial:", err)
	}

	log.Println("create connection websocket")

	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		log.Println("running reader")
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read socket error:", err)
				return
			}
			log.Println("message socket", message)
		}
	}()

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	ping := akun.createPing()

	log.Println("sending ping")

	err = c.WriteMessage(websocket.BinaryMessage, ping)
	if err != nil {
		log.Println("write:", err)
		return
	}

	for {
		select {
		case <-done:
			return
		}
	}
}
