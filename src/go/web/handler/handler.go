package handler

import (
	"net/http"
    "bytes"
    "encoding/json"
    "github.com/jabaraster/webtool"
    "fmt"
)

type commandRequest struct {
    Name string `json:"name"`
    Parameters map[string]string `json:"parameters"`
}


func StartSessionHandler(w http.ResponseWriter, r *http.Request) {

    var req commandRequest
    mustDecodeRequestJson(r, &req)

    fmt.Println("["+req.Name+"]")

    switch req.Name {
        case "camera.startSession":
            startSession(&req, w)
        case "camera.takePicture":
            takePicture(&req, w)
        case "camera.getImage":
            getImage(&req, w)
        case "camera.closeSession":
            closeSession(&req, w)
        default:
            panic(req.Name + " is invalid command.")
    }
}

func startSession(r *commandRequest, w http.ResponseWriter) {
    type Res struct {
        Name string `json:"name"`
        State string `json:"state"`
        Results map[string]interface{} `json:"results"`
    }
    res := Res{
        Name: r.Name,
        State: "done",
        Results: map[string]interface{}{
            "sessionId": "SID_0001",
            "timeout": 180,
        },
    }
    webtool.WriteJsonResponse(res, w)
}

func takePicture(r *commandRequest, w http.ResponseWriter) {
    fmt.Println(r)
    type Res struct {
        Name string `json:"name"`
        State string `json:"state"`
        Id int64 `json:"id"`
        Progress map[string]int64 `json:"progress"`
    }
    res := Res{
        Name: r.Name,
        State: "inProgress",
        Id: 1,
        Progress: map[string]int64 {
            "completion": 0,
        },
    }
    webtool.WriteJsonResponse(res, w)
}

func getImage(r *commandRequest, w http.ResponseWriter) {
    fmt.Println(r)
}

func closeSession(r *commandRequest, w http.ResponseWriter) {
    fmt.Println(r)
    // nop
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
