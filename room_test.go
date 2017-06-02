package twilio

import "testing"

const (
	apiKey    = "SK7e92aab55f2a7f09af7a14493c2d791c"
	apiSecret = "re0W7vK4ZNT2tPxetvHY8A86nfTvq6gD"
)

func TestGetRoom(t *testing.T) {
	tw := NewTwilio(apiKey, apiSecret, nil)
	tw.EnableDebug()

	// Exist room
	room, err := tw.GetRoom("DailyStandup")
	if err != nil {
		t.Errorf("Get exists room failed: %v", err)
		t.Fail()
	}

	t.Logf("Room: %+v", room)
}
