package sentinel
import (
    "path"
    "os"
    "log"
    "net/http"
    "io"
)

type torrent struct {
    Id              int     `json:"id"`
    Hash            string  `json:"hash"`
    Name            string  `json:"name"`
    Filename        string  `json:"filename"`
    Freeleech       string  `json:"freeleech"`
    Added           string  `json:"added"`
    Comments        uint    `json:"comments"`
    Numfiles        uint    `json:"numfiles"`
    Leechers        uint    `json:"leechers"`
    Seeders         uint    `json:"seeders"`
    Times_completed uint    `json:"times_completed"`
    Size            uint    `json:"size"`
    Utadded         uint    `json:"utadded"`
    Type_category   uint    `json:"type_category"`
    Type_codec      uint    `json:"type_codec"`
    Type_medium     uint    `json:"type_medium"`
    TypeOrigin     uint    `json:"type_origin"`
    TorrentURL      string  `json:"-"`
    Additional      params  `json:"-"`
}

type torrents []torrent

func (data torrents) Len() int {
    return len(data)
}

func (data torrents) Swap(a, b int) {
    data[a], data[b] = data[b], data[a]
}

func (data torrents) Less(a, b int) bool {
    return data[a].Added > data[b].Added //Reverse the order
}

func (tor torrent) Download() bool {
    output := path.Join(Config.Output, tor.Filename)

    fd, err := os.Create(output)
    if err != nil {
        log.Println(err.Error())
        return false
    }
    defer fd.Close()

    resp, err := http.Get(tor.TorrentURL)
    if err != nil {
        log.Println(err.Error())
        return false
    }
    defer resp.Body.Close()

    _, err = io.Copy(fd, resp.Body)
    if err != nil {
        log.Println(err.Error())
        return false
    }
    return true
}