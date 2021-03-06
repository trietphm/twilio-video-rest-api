package twilio

type Error struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   int    `json:"status"`
}

func (e Error) Error() string {
	return e.Message
}

func ParseTwilioError(err error) (Error, error) {
	te, ok := err.(Error)
	if ok {
		return te, nil
	}

	return te, err
}
