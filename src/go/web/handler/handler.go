package handler

import (
	"github.com/jabaraster/webtool"
	"net/http"
    "bytes"
    "encoding/json"
)

func GetAllPropertiesHandler(w http.ResponseWriter, r *http.Request) {
    type D struct {
        Name string
        Parameters map[string]string
    }
    var data D
    mustDecodeRequestJson(r, &data)
}

func mustDecodeRequestJson(r *http.Request, v interface{}) {
    buffer := new(bytes.Buffer)
    _, bre := buffer.ReadFrom(r.Body)
    if bre != nil {
        panic(bre)
    }

    err := json.Unmarshal(buffer.Bytes(), v)
    if err != nil {
        panic(err)
    }
}
