package globe

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type FullRequest struct {
	Translation string `json:"translation"`
}

type StringRequest struct {
	Key string         `json:"key"`
	Translation string `json:"translation"`
}

type FullResponse struct {
	Key string    `json:"key"`
	String string `json:"string"`
}

type StringResponse struct {
	String string `json:"string"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func StringRequestHandler(g *GlobeDB) func (w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		req := &StringRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(req)
		var resp []byte
		var respStr string
		
		fmt.Printf("Got request: %#v\n", req)

		if err != nil {
			goto error
		}

		respStr, err = g.Lookup(req.Key, req.Translation)
		if err != nil {
			goto error
		}

		resp, err = json.Marshal(&StringResponse{respStr})
		fmt.Fprintf(w, "%s", resp)
		return

	error:
		fmt.Printf("Error: %s\n", err)
		reply, err := json.Marshal(&ErrorResponse{err.Error()})
		if err != nil {
			fmt.Printf("Failed to marshal error: %s\n", err)
			http.Error(w, `{"Error": "Unknown"}`, 500)
			return
		}
		fmt.Fprintf(w, "%s", reply)
		return
		
	}
}

func FullTranslationRequestHandler(g *GlobeDB) func (w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		req := &FullRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(req)
		var resp []byte
		var all map[string]string
		var respObjects []FullResponse
		
		if err != nil {
			goto error
		}

		all = g.LookupAll(req.Translation)
		respObjects = make([]FullResponse, 0)
		for k, v := range all {
			respObjects = append(respObjects, FullResponse{k, v})
		}

		resp, err = json.Marshal(respObjects)
		if err != nil {
			goto error
		}
		fmt.Fprintf(w, "%s", resp)
		return

	error:
		fmt.Printf("Error: %s\n", err)
		reply, err := json.Marshal(&ErrorResponse{err.Error()})
		if err != nil {
			fmt.Printf("Failed to marshal error: %s\n", err)
			http.Error(w, `{"Error": "Unknown"}`, 500)
			return
		}
		fmt.Fprintf(w, "%s", reply)
		return
	}
}
