//https://www.twilio.com/docs/api/video/rooms-resource#resource-properties

package twilio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/kevinburke/rest"
)

const (
	basePath = "https://video.twilio.com/v1/"

	RoomTypePeerToPeer = "peer-too-peer"
	RoomTypeGroup      = "group"

	RoomStatusInProgrcess = "in-progress"
	RoomStatusFailed      = "failed"
	RoomStatusCompleted   = "completed"

	PeerToPeerMaxParticipants = 10

	TimeFormat = "2006-01-02T15:04:05Z" //ISO 8601

	defaultTimeout = 30 * time.Second
)

type twilio struct {
	ApiKey    string
	ApiSecret string
	debug     bool
}

type RoomParam struct {
	Type                        string
	EnableTurn                  bool
	UniqueName                  string
	StatusCallback              string
	StatusCallbackMethod        string
	RecordParticipantsOnConnect string
	MaxParticipants             int
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

var client *http.Client

func init() {
	client = &http.Client{
		Timeout:   defaultTimeout,
		Transport: rest.DefaultTransport,
	}
}

func NewTwilio(apiKey, apiSecret string, c *http.Client) twilio {
	if c != nil {
		client = c
	}
	return twilio{
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
	}
}

func (t *twilio) EnableDebug() {
	t.debug = true
}

func (t *twilio) DisableDebug() {
	t.debug = false
}

func (t *twilio) GetRoom(roomName string) (room Room, err error) {
	var body []byte
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

	//TODO handle status

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &room)
	return
}

func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n\n", data)
	} else {
		fmt.Printf("[ERROR]\n %s\n\n", err)
	}
}
