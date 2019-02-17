package globe

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type StringResponse struct {
	String string `json:"string"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func RequestHandler(g *GlobeDB) func (w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		req := &StringRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(req)

		fmt.Printf("Got request: %#v\n", req)

		if err != nil {
			fmt.Printf("Error: %s\n", err)
			r, err := json.Marshal(&ErrorResponse{err.Error()})
			if err != nil {
				http.Error(w, `{"Error": "Unknown"}`, 500)
				return
			}
			fmt.Fprintf(w, "%s", r)
			return
		}

		respStr, err := g.Lookup(req.StringName, req.Translation)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			r, err := json.Marshal(&ErrorResponse{err.Error()})
			if err != nil {
				http.Error(w, `{"Error": "Unknown"}`, 500)
				return
			}
			fmt.Fprintf(w, "%s", r)
			return
		}

		resp, err := json.Marshal(&StringResponse{respStr})
		fmt.Fprintf(w, "%s", resp)
	}
}
