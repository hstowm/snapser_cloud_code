package response

import "encoding/json"

type ResMessage struct {
	Code       int              `json:"code"`
	Message    string           `json:"message"`
	Currencies map[string]int32 `json:"currencies"`
}

func (res *ResMessage) toString() string {
	ret, err := json.Marshal(res)
	if err != nil {
		return ""
	}
	return string(ret)
}
