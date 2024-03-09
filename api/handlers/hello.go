package handlers

import (
	"fmt"
	"net/http"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request){
	fmt.Fprint(writer, "Hello Web")
}
