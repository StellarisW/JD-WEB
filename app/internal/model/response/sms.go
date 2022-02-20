package response

type SmsRes struct {
	Sign    string `json:"sign"`
	SmsCode string `json:"smsCode"`
}
