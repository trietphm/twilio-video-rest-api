//https://www.twilio.com/docs/api/video/rooms-resource#resource-properties

package twilio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

const (
	basePath = "https://video.twilio.com/v1/"

	RoomTypePeerToPeer = "peer-to-peer"
	RoomTypeGroup      = "group"

	RoomStatusInProgrcess = "in-progress"
	RoomStatusFailed      = "failed"
	RoomStatusCompleted   = "completed"

	PeerToPeerMaxParticipants = 10

	TimeFormat = "2006-01-02T15:04:05Z" //ISO 8601

	defaultTimeout = 30 * time.Second
)

type Twilio struct {
	ApiKey    string
	ApiSecret string
	debug     bool
}

type Room struct {
	Sid                         string     `json:"sid"`
	Status                      string     `json:"status"`
	DateCreated                 time.Time  `json:"date_created"`
	DateUpdated                 time.Time  `json:"date_updated"`
	AccountSid                  string     `json:"account_sid"`
	Type                        string     `json:"type"`
	EnableTurn                  bool       `json:"enable_turn"`
	UniqueName                  string     `json:"unique_name"`
	StatusCallback              *string    `json:"status_callback"`
	StatusCallbackMethod        string     `json:"status_callback_method"`
	EndTime                     *time.Time `json:"end_time"`
	Duration                    *int       `json:"duratin"`
	MaxParticipants             int        `json:"max_participants"`
	RecordParticipantsOnConnect bool       `json:"record_participants_on_connect"`
	Url                         string     `json:"url"`
	Links                       Link       `json:"links"`
}

type Link struct {
	Recordings string `json:"recordings"`
}

type ListRoom struct {
	Meta  Meta   `json:"meta"`
	Rooms []Room `json:"rooms"`
}

type Meta struct {
	Page            int     `json:"page"`
	PageSize        int     `json:"page_size"`
	FirstPageUrl    string  `json:"first_page_url"`
	PreviousPageUrl *string `json:"previous_page_url"`
	Url             string  `json:"url"`
	NextPageUrl     *string `json:"next_page_url"`
	Key             string  `json:"key"`
}

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: defaultTimeout,
	}
}

func NewTwilio(apiKey, apiSecret string, c *http.Client) Twilio {
	if c != nil {
		client = c
	}
	return Twilio{
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
	}
}

func (t *Twilio) EnableDebug() {
	t.debug = true
}

func (t *Twilio) DisableDebug() {
	t.debug = false
}

func (t *Twilio) GetRoom(roomName string) (room Room, err error) {
	var response *http.Response
	var request *http.Request

	url := fmt.Sprintf(basePath+"Rooms/%s", roomName)
	request, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	request.SetBasicAuth(t.ApiKey, t.ApiSecret)

	// Dump request
	if t.debug {
		debug(httputil.DumpRequestOut(request, false))
	}

	response, err = client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if t.debug {
		debug(httputil.DumpResponse(response, true))
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	if response.StatusCode != 200 && response.StatusCode != 201 && response.StatusCode != 202 {
		var resErr Error
		err = json.Unmarshal(body, &resErr)
		if err != nil {
			return
		}

		return room, resErr
	}

	err = json.Unmarshal(body, &room)
	return
}

func (t *Twilio) CreateRoom(param roomParam) (room Room, err error) {
	var response *http.Response
	var request *http.Request

	url := basePath + "Rooms"
	data := param.toURLEncoded()
	requestBody := strings.NewReader(data.Encode())
	request, err = http.NewRequest("POST", url, requestBody)
	if err != nil {
		return
	}
	request.SetBasicAuth(t.ApiKey, t.ApiSecret)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	// Dump request
	if t.debug {
		fmt.Println("[DEBUG][RequestBody]")
		debug(httputil.DumpRequestOut(request, true))
	}

	response, err = client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if t.debug {
		fmt.Println("[DEBUG][ResponseBody]")
		debug(httputil.DumpResponse(response, true))
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	if response.StatusCode != 200 && response.StatusCode != 201 && response.StatusCode != 202 {
		var resErr Error
		err = json.Unmarshal(body, &resErr)
		if err != nil {
			return
		}

		return room, resErr
	}

	err = json.Unmarshal(body, &room)
	return
}

func (t *Twilio) CompleteRoom(roomName string) (room Room, err error) {
	var response *http.Response
	var request *http.Request

	link := fmt.Sprintf(basePath+"Rooms/%s", roomName)
	data := url.Values{}
	data.Set("Status", RoomStatusCompleted)
	requestBody := strings.NewReader(data.Encode())
	request, err = http.NewRequest("POST", link, requestBody)
	if err != nil {
		return
	}
	request.SetBasicAuth(t.ApiKey, t.ApiSecret)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	// Dump request
	if t.debug {
		fmt.Println("[DEBUG][RequestBody]")
		debug(httputil.DumpRequestOut(request, true))
	}

	response, err = client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if t.debug {
		fmt.Println("[DEBUG][ResponseBody]")
		debug(httputil.DumpResponse(response, true))
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	if response.StatusCode != 200 && response.StatusCode != 201 && response.StatusCode != 202 {
		var resErr Error
		err = json.Unmarshal(body, &resErr)
		if err != nil {
			return
		}

		return room, resErr
	}

	err = json.Unmarshal(body, &room)
	return
}

func (t *Twilio) GetListRooms(param listRoomParam) (listRoom ListRoom, err error) {
	var response *http.Response
	var request *http.Request

	url := basePath + "Rooms"
	request, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	request.SetBasicAuth(t.ApiKey, t.ApiSecret)
	queryParam := param.toQueryParam()
	request.URL.RawQuery = queryParam.Encode()

	// Dump request
	if t.debug {
		debug(httputil.DumpRequestOut(request, false))
	}

	response, err = client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if t.debug {
		debug(httputil.DumpResponse(response, true))
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	if response.StatusCode != 200 && response.StatusCode != 201 && response.StatusCode != 202 {
		var resErr Error
		err = json.Unmarshal(body, &resErr)
		if err != nil {
			return
		}

		return listRoom, resErr
	}

	err = json.Unmarshal(body, &listRoom)
	return
}

func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n\n", data)
	} else {
		fmt.Printf("[ERROR]\n %s\n\n", err)
	}
}
