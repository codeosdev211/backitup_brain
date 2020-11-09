package commons

import (
    http "net/http"
    "encoding/json"
    models "../models"
)

func SendErrorResponse(res *http.ResponseWriter, msg string) {
    json.NewEncoder(*res).Encode(models.Response{1, msg, nil})
}
