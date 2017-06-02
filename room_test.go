package twilio

import "testing"

const (
	apiKey    = ""
	apiSecret = ""
)

func TestGetRoom(t *testing.T) {
	tw := NewTwilio(apiKey, apiSecret, nil)

	// Exist room
	room, err := tw.GetRoom("DailyStandup")
	if err != nil {
		t.Errorf("Get exists room failed: %v", err)
		t.Fail()
	}

	t.Logf("Room: %+v", room)
}

func TestNotFoundRoom(t *testing.T) {
	tw := NewTwilio(apiKey, apiSecret, nil)

	// Exist room
	room, err := tw.GetRoom("THISISNOTFOUNDROOM")
	if err == nil {
		t.Errorf("Get not found room failed, room: %v", room)
		t.Fail()
	}

	notFoundError, ok := err.(Error)
	if !ok {
		t.Errorf("Error is not twilio error: %v", err)
		t.Fail()
	}

	if notFoundError.Status != 404 {
		t.Errorf("Error is not not found: %v", err)
		t.Fail()
	}

	t.Logf("Room is not found, err: %v", notFoundError)
}

func TestAuthorizationError(t *testing.T) {
	tw := NewTwilio("abcd", "abcd", nil)
	tw.EnableDebug()

	// Exist room
	room, err := tw.GetRoom("RANDOM_ROOM")
	if err == nil {
		t.Errorf("Get not found room failed, room: %v", room)
		t.Fail()
	}

	authError, ok := err.(Error)
	if !ok {
		t.Errorf("Error is not twilio error: %v", err)
		t.Fail()
	}

	if authError.Status != 401 {
		t.Errorf("Error is not Authenticate error: %v", err)
		t.Fail()
	}

	t.Logf("Authencation fail, err: %v", authError)
}
