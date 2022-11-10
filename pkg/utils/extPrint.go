package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func PrintlnWithTimeShtamp(text string) {
	fmt.Println(time.Now().Format(time.RFC3339), " :", text)
}
