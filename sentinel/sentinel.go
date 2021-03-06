package sentinel

import (
    "path"
    "regexp"
    "io/ioutil"
    "encoding/json"
    "os/user"
)

type config struct {
    Username    string          `json:"username"`
    Passkey     string          `json:"passkey"`
    Output      string          `json:"output`
    IRC         *irc_config     `json:"irc"`
}
type commands []string
type irc_config struct {
    Host        string      `json:"host"`
    Port        int         `json:"port"`
    Password    string      `json:"password"`
    SSL         bool        `json:"ssl"`
    SSL_SKIP    bool        `json:"insecure_ssl_verify_skip"`
    Nick        string      `json:"nick"`
    Stalk       []string    `json:"stalk"`
    Channels    []string    `json:"channels"`
    Commands    []commands  `json:"commands"`
    StalkSet    set         `json:"-"`
    ChannelSet  set         `json:"-"`
}

func (c *config) BuildChannelSet() {
    channel_set := make(set)
    for _, channel := range c.IRC.Channels {
        channel_set[channel] = true
    }
    c.IRC.ChannelSet = channel_set
}

func (c *config) BuildStalkSet() {
    stalker := make(set)
    for _, stalk := range c.IRC.Stalk {
        stalker[stalk] = true
    }

    c.IRC.StalkSet = stalker
}

var re *regexp.Regexp
var Config config
var Filter filter

const regex string = "^New Torrent: (?P<torrent>.*?(?P<resolution>\\d+(?:p|i)).*?(?P<source>(?:BluRay|Blu-Ray|HDTV|WEB-DL|Capture)).*?) - " +
"Type: (?P<category>Music|Documentary|Misc/Demo|Movie|Sport|TV|XXX) " +
"\\((?P<codec>(?:H.264|MPEG-2|VC-1|XviD)), (?P<medium>(?:Blu-ray/HD DVD|Capture|Encode|Remux|WEB-DL))\\) " +
"((?P<origin>Internal|)! )?- Uploaded by: (?P<uploader>.*?)$"

var categories map[string]int = map[string]int{"Movie": 1, "TV": 2, "Documentary": 3, "Music": 4, "Sport": 5, "Audio Track": 6, "XXX": 7, "Misc/Demo": 8}
var codec map[string]int = map[string]int{"H.264": 1, "MPEG-2": 2, "VC-1": 3, "XviD": 4}
var mediums map[string]int = map[string]int{"Blu-ray/HD DVD": 1, "Encode": 3, "Capture": 4, "Remux": 5, "WEB-DL": 6}
var origin map[string]int = map[string]int{"Internal": 1}

type params map[string]string
type set map[string]bool

func init() {
    u, _ := user.Current()
    c, _ := ioutil.ReadFile(path.Join(u.HomeDir, ".beaconrc"))
    f, _ := ioutil.ReadFile(path.Join(u.HomeDir, ".filter.beacon"))

    json.Unmarshal(c, &Config)
    json.Unmarshal(f, &Filter)

    Config.BuildStalkSet()
    Config.BuildChannelSet()
}