package sentinel

import (
    "regexp"
    "io/ioutil"
    "encoding/json"
)

type config struct {
    Username    string          `json:"username"`
    Passkey     string          `json:"passkey"`
    Stalk       []string        `json:"stalk"`
    Channels    []string        `json:"channels"`
    Output      string          `json:"output`
    ZNC         *znc_config     `json:"znc"`
    IRC         *irc_config     `json:"irc"`
    StalkSet    set             `json:"-"`
    ChannelSet  set             `json:"-"`
}

type proto struct {
    Host        string          `json:"host"`
    Port        int             `json:"port"`
}
type znc_config struct {
    proto
    Password    string      `json:"password"`
    SSL         bool        `json:"ssl"`
    SSL_SKIP    bool        `json:"insecure_ssl_verify_skip"`
}
type irc_config struct {
    proto
    Nick        string      `json:"nick"`
}

func (c *config) BuildChannelSet() {
    c.ChannelSet = make(set)
    for _, channel := range c.Channels {
        c.ChannelSet[channel] = true
    }
}

func (c *config) BuildStalkSet() {
    c.StalkSet = make(set)
    for _, stalk := range c.Stalk {
        c.StalkSet[stalk] = true
    }
}

var re *regexp.Regexp
var Config config
var Filter filter

const regex string = "^New Torrent: (?P<torrent>.*?(?P<resolution>\\d+(?:p|i)).*?(?P<source>(?:BluRay|Blu-Ray|HDTV)).*?) - " +
"Type: (?P<category>Audio Track|Documentary|Misc/Demo|Movie|Music|Sport|TV|XXX) " +
"\\((?P<codec>(?:H.264|MPEG-2|VC-1|XviD)), (?P<medium>(?:Blu-ray/HD DVD|Capture|Encode|Remux|WEB-DL))\\) " +
"((?P<origin>Internal|)! )?- Uploaded by: (?P<uploader>.*?)$"

var categories map[string]int = map[string]int{"Movie": 1, "TV": 2, "Documentary": 3, "Music": 4, "Sport": 5, "Audio Track": 6, "XXX": 7, "Misc/Demo": 8}
var codec map[string]int = map[string]int{"H.264": 1, "MPEG-2": 2, "VC-1": 3, "XviD": 4}
var mediums map[string]int = map[string]int{"Blu-ray/HD DVD": 1, "Encode": 3, "Capture": 4, "Remux": 5, "WEB-DL": 6}
var origin map[string]int = map[string]int{"Internal": 1}

type params map[string]string
type set map[string]bool
func init() {
    c, _ := ioutil.ReadFile("~/.beaconrc")
    json.Unmarshal(c, &Config)

    f, _ := ioutil.ReadFile("~/.filter.beacon")
    json.Unmarshal(f, &Filter)
    Filter.PopulateSets()
}