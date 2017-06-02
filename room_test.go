package twilio

import "testing"

const (
	apiKey    = ""
	apiSecret = ""
)

func TestCreateRoomSuccess(t *testing.T) {
	tw := NewTwilio(apiKey, apiSecret, nil)
	param := roomParam{
		Type:                        RoomTypePeerToPeer,
		EnableTurn:                  true,
		UniqueName:                  "TestRoom",
		StatusCallback:              "http://twilio.com",
		StatusCallbackMethod:        "GET",
		RecordParticipantsOnConnect: false,
		MaxParticipants:             2,
	}

	room, err := tw.CreateRoom(param)
	if err != nil {
		te, err := ParseTwilioError(err)
		if err != nil {
			t.Errorf("Error is not twilio error: %v", err)
			t.Fail()
		}

		t.Errorf("Create room failed: %+v", te)
		t.Fail()
		return
	}

	t.Logf("Room: %+v", room)
}

func TestCreateRoomFail(t *testing.T) {
	tw := NewTwilio(apiKey, apiSecret, nil)
	param := roomParam{
		Type:                        "invalidType",
		EnableTurn:                  true,
		UniqueName:                  "TestRoom",
		StatusCallback:              "http://twilio.com",
		StatusCallbackMethod:        "GET",
		RecordParticipantsOnConnect: false,
		MaxParticipants:             2,
	}

	room, err := tw.CreateRoom(param)
	if err == nil {
		t.Errorf("Create room success, room: %v", room)
		t.Fail()
		return
	}

	badParam, err := ParseTwilioError(err)
	if err != nil {
		t.Errorf("Error is not twilio error: %v", err)
		t.Fail()
		return
	}

	if badParam.Status != 400 {
		t.Errorf("Error is not bad param: %v", err)
		t.Fail()
		return
	}

	t.Logf("Success error: %+v", err)
}

func TestGetRoom(t *testing.T) {
	tw := NewTwilio(apiKey, apiSecret, nil)

	// Exist room
	room, err := tw.GetRoom("TestRoom")
	if err != nil {
		te, err := ParseTwilioError(err)
		if err != nil {
			t.Errorf("Error is not twilio error: %v", err)
			t.Fail()
		}

		t.Errorf("Create room failed: %v", te)
		t.Fail()
		return
	}

	t.Logf("Room: %+v", room)
}

func TestGetNotFoundRoom(t *testing.T) {
	tw := NewTwilio(apiKey, apiSecret, nil)

	room, err := tw.GetRoom("THISISNOTFOUNDROOM")
	if err == nil {
		t.Errorf("Get not found room failed, room: %v", room)
		t.Fail()
	}

	notFoundError, err := ParseTwilioError(err)
	if err != nil {
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

	room, err := tw.GetRoom("RANDOM_ROOM")
	if err == nil {
		t.Errorf("Get not found room failed, room: %v", room)
		t.Fail()
	}

	authError, err := ParseTwilioError(err)
	if err != nil {
		t.Errorf("Error is not twilio error: %v", err)
		t.Fail()
	}

	if authError.Status != 401 {
		t.Errorf("Error is not Authenticate error: %v", err)
		t.Fail()
	}

	t.Logf("Authencation fail, err: %v", authError)
}

func TestCompleteRoom(t *testing.T) {
	tw := NewTwilio(apiKey, apiSecret, nil)
	tw.EnableDebug()

	// Exist room
	room, err := tw.CompleteRoom("TestRoom")
	if err != nil {
		te, err := ParseTwilioError(err)
		if err != nil {
			t.Errorf("Error is not twilio error: %v", err)
			t.Fail()
		}

		t.Errorf("Complete room failed: %v", te)
		t.Fail()
		return
	}

	t.Logf("Room completed: %+v", room)
}
