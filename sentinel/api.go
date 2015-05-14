package sentinel
import (
    "io/ioutil"
    "net/http"
    "bytes"
    "sort"
    "fmt"
    "net/url"
)

func Fetch(announce string) torrent {
    results := announce_matcher(announce)

    payload := init_payload(results)

    post, _ := http.Post("https://hdbits.org/api/torrents", "application/json", bytes.NewReader(payload.ToBytes()))
    defer post.Body.Close()

    response, _ := ioutil.ReadAll(post.Body)

    resp := build_response(response)

    sort.Sort(torrents(resp.Data))
    torrent := resp.Data[0]
    encoded_filename := url.QueryEscape(torrent.Filename)
    torrent.TorrentURL = fmt.Sprintf("https://hdbits.org/download.php/%s?id=%d&passkey=%s", encoded_filename, torrent.Id, Config.Passkey)
    torrent.Additional = results
    return torrent
}