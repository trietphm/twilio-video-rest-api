package twilio

import (
	"net/url"
	"strconv"
	"time"
)

type roomParam struct {
	Type                        string
	EnableTurn                  bool
	UniqueName                  string
	StatusCallback              string
	StatusCallbackMethod        string
	RecordParticipantsOnConnect bool
	MaxParticipants             int
}

type listRoomParam struct {
	Status            string
	DateCreatedAfter  *time.Time
	DateCreatedBefore *time.Time
	UniqueName        string
}

func NewRoomParam() roomParam {
	return roomParam{
		Type:                        "",
		EnableTurn:                  true,
		UniqueName:                  "",
		StatusCallback:              "",
		StatusCallbackMethod:        "POST",
		RecordParticipantsOnConnect: false,
		MaxParticipants:             50,
	}
}

func NewListRoomParam() listRoomParam {
	return listRoomParam{
		Status:            "",
		DateCreatedAfter:  nil,
		DateCreatedBefore: nil,
		UniqueName:        "",
	}
}

func (r roomParam) toURLEncoded() url.Values {
	data := url.Values{}
	data.Set("Type", r.Type)
	data.Set("EnableTurn", strconv.FormatBool(r.EnableTurn))
	data.Set("UniqueName", r.UniqueName)
	data.Set("StatusCallback", r.StatusCallback)
	data.Set("StatusCallbackMethod", r.StatusCallbackMethod)
	data.Set("RecordParticipantsOnConnect", strconv.FormatBool(r.RecordParticipantsOnConnect))
	data.Set("MaxParticipants", strconv.Itoa(r.MaxParticipants))
	return data
}

func (p listRoomParam) toQueryParam() url.Values {
	dateLayout := "2006-01-02"
	data := url.Values{}
	if p.Status != "" {
		data.Set("Status", p.Status)
	}
	if p.DateCreatedBefore != nil {
		data.Set("DateCreatedBefore", p.DateCreatedBefore.Format(dateLayout))
	}
	if p.DateCreatedAfter != nil {
		data.Set("DateCreatedAfter", p.DateCreatedAfter.Format(dateLayout))
	}

	if p.UniqueName != "" {
		data.Set("UniqueName", p.UniqueName)
	}
	return data
}
