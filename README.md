Twilio REST API Tool
=========

[![GoDoc](https://godoc.org/github.com/trietphm/twilio-video-rest-api?status.svg)](http://godoc.org/github.com/trietphm/twilio-video-rest-api) [![Build Status](https://travis-ci.org/trietphm/twilio-video-rest-api.svg?branch=master)](https://travis-ci.org/trietphm/twilio-video-rest-api)

Easier to access to Twilio Programmable Video REST API.
Twilio Document: https://www.twilio.com/docs/api/video/rest 

## Install 

```
go get github.com/trietphm/twilio-video-rest-api
```

## Usage
### Twilio REST
First, get your own API key & secrect from [Twilio Console](https://www.twilio.com/console/video/dev-tools/api-keys)

Create new Twilio 

```
// Pass your httpClient or leave it nil, then default http client will be used
tw := twilio.NewTwilio(ApiKey, ApiSecret, httpClient) 
```

Get a room 

```
room, err := tw.GetRoom("MyRoom")
if err != nil {
	panic(err)
	/* To get original Twilio error, parse it to TwilioError
	te, err := ParseTwilioError(err)
	if err != nil {
		panic("Parse twilio error fail")
	}
	panic (te)
	*/
}
fmt.Printf("%+v\n", room)
```

### Debug tool 
You can log all information of your request/response to Twilio REST api. 
To enable debug:

```
tw.EnableDebug()
```

And disable it 

```
tw.DisableDebug()
```

Debug tool will be disable by default

### Supported

- [x] Create Room
- [x] Get Room
- [x] Complete Room
- [x] Get list Room
- [x] Debug tool

