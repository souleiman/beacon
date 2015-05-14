package sentinel
import (
    "encoding/json"
)

type message struct {
    Status  int     `json:"status"`
    Message string  `json:"message"`
}

type response struct {
    message
    Data    torrents `json:"data"`
}

func build_response(resp []byte) response {
    var r response
    json.Unmarshal(resp, &r)

    return r
}