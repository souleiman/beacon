package sentinel
import "encoding/json"

type auth struct {
    Username    string  `json:"username"`
    Passkey     string  `json:"passkey"`
}

type torrent_payload struct {
    auth
    Search      string  `json:"search"`
    Category    []int   `json:"category"`
    Medium      []int   `json:"medium"`
    Codec       []int   `json:"codec"`
    Origin      []int   `json:"origin"`
}

func init_payload(torrent map[string]string) torrent_payload {
    payload := torrent_payload{}
    payload.Username = Config.Username
    payload.Passkey = Config.Passkey
    payload.Search = torrent["torrent"]
    payload.Category = []int{categories[torrent["category"]]}
    payload.Medium = []int{mediums[torrent["medium"]]}
    payload.Codec = []int{codec[torrent["codec"]]}
    payload.Origin = []int{origin[torrent["origin"]]}

    return payload
}

func (payload torrent_payload) ToBytes() (b []byte) {
    b, _ = json.Marshal(payload)
    return
}