package utils

// Api response codes
const (
	ResponseFailed    = int32(1)
	ResponseOk        = int32(0)
	MaxRestAPIPayload = 1073741824
)

// APIResp ...
type APIResp struct {
	ResponseCode        int32  `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
}
