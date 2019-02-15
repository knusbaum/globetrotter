package globe

import (
	"fmt"
	"net/http"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request form: %#v\n", r.Form)
}
