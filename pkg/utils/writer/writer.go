package writer

import (
	"encoding/json"
	"log"
	"net/http"
)

func HTTPResponseWriter(resp http.ResponseWriter, err error, body interface{}) {
	resp.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(resp).Encode(body)

	if err != nil {
		log.Printf("couldn't encode the value[%v] into json: %v", body, err)
	}
}
