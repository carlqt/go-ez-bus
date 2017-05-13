package render

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, h interface{}) {
	resp, err := json.Marshal(h)
	if err != nil {
		log.Println(err)
	}

	w.Write(resp)
}
