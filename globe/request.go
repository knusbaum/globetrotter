package globe

import (
	"encoding/json"
)

type StringRequest struct {
	StringName string  `json:"stringName"`
	Translation string `json:"translation"`
}

func ParseStringRequest(req string) (*StringRequest, error) {
	strReq := &StringRequest{}
	err := json.Unmarshal([]byte(req), strReq)
	if err != nil {
		return nil, err
	}

	return strReq, nil
}

func LookupJson(g *GlobeDB, req string) (string, error) {
	strReq, err := ParseStringRequest(req)
	if err != nil {
		return "", err
	}

	return g.Lookup(strReq.StringName, strReq.Translation)
}


