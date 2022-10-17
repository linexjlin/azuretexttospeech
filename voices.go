package azuretexttospeech

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/proxy"
	gtls "gopkg.in/h2non/gentleman.v2/plugins/tls"
)

// voiceListAPI is the source for supported voice list to region mapping
// See: https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/rest-text-to-speech#regions-and-endpoints
const voiceListAPI = "https://%s.tts.speech.microsoft.com/cognitiveservices/voices/list"

//go:generate enumer -type=voiceType -linecomment -json
type voiceType int

const (
	voiceStandard voiceType = iota // Standard
	voiceNeural                    // Neural
)

type regionVoiceListResponse struct {
	Name            string    `json:"Name"`
	ShortName       string    `json:"ShortName"`
	Gender          Gender    `json:"Gender"`
	Locale          string    `json:"Locale"`
	SampleRateHertz string    `json:"SampleRateHertz"`
	VoiceType       voiceType `json:"VoiceType"`
}

// supportedVoices represents the key used within the `localeToGender` map.
type supportedVoices struct {
	Gender Gender
	Locale string
}

type RegionVoiceMap map[supportedVoices]string

func (az *AzureCSTextToSpeech) buildVoiceToRegionMap() (RegionVoiceMap, error) {

	v, err := az.fetchVoiceList()
	if err != nil {
		return nil, err
	}

	m := make(map[supportedVoices]string)
	for _, x := range v {
		if x.VoiceType == voiceNeural {
			m[supportedVoices{Gender: x.Gender, Locale: x.Locale}] = x.ShortName
		}
	}
	return m, err
}

func (az *AzureCSTextToSpeech) fetchVoiceList() ([]regionVoiceListResponse, error) {

	// Create a new client
	cli := gentleman.New()

	// Define base URL
	cli.URL(az.voiceServiceListURL)

	if az.HttpProxy != "" {

		servers := map[string]string{"http": az.HttpProxy, "https": az.HttpProxy}
		cli.Use(proxy.Set(servers))
	}

	cli.Use(gtls.Config(&tls.Config{InsecureSkipVerify: true}))

	// Create a new request based on the current client
	req := cli.Request()

	// Set a new header field
	req.SetHeader("Authorization", "Bearer "+az.accessToken)

	// Perform the request
	res, err := req.Send()
	if err != nil {
		return nil, err
	}
	if !res.Ok {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusOK:
		var r []regionVoiceListResponse
		if err := json.Unmarshal(res.Bytes(), &r); err != nil {
			return nil, fmt.Errorf("unable to decode voice list response body, %v", err)
		}
		return r, nil
	case http.StatusBadRequest:
		return nil, fmt.Errorf("%d - A required parameter is missing, empty, or null. Or, the value passed to either a required or optional parameter is invalid. A common issue is a header that is too long", res.StatusCode)
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("%d - The request is not authorized. Check to make sure your subscription key or token is valid and in the correct region", res.StatusCode)
	case http.StatusTooManyRequests:
		return nil, fmt.Errorf("%d - You have exceeded the quota or rate of requests allowed for your subscription", res.StatusCode)
	case http.StatusBadGateway:
		return nil, fmt.Errorf("%d - Network or server-side issue. May also indicate invalid headers", res.StatusCode)
	}
	return nil, fmt.Errorf("%d - unexpected response code from voice list API", res.StatusCode)
}
