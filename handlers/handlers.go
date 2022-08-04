package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func RandomsHandler(w http.ResponseWriter, r *http.Request) {

	writeString, err := io.WriteString(w, "Hello World")
	if err != nil {
		return
	} else {
		fmt.Printf("%v", writeString)
	}

}
