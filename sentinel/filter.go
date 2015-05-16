package sentinel

type filter struct {
    Internal    bool    `json:"always_dl_internal"`
    MaxSeeders  int     `json:"maximum_seeders"`
    Resolutions set     `json:"allowed_resolution"`
    Mediums     set     `json:"allowed_medium"`
    Codecs      set     `json:"allowed_codec"`
    Category    set     `json:"allowed_category"`
    Uploader    set     `json:"allowed_uploader"`
}

func (f *filter) Valid(tor torrent) bool {
    if f.Internal && int(tor.TypeOrigin) == 1{
        return true
    }

    if f.MaxSeeders > 1 && int(tor.Seeders) >= f.MaxSeeders {
        return false
    }

    add := tor.Additional
   	allow := f.Resolutions[add["resolution"]] && f.Mediums[add["medium"]] && f.Codecs[add["codec"]] && f.Category[add["category"]]
   	return (len(f.Uploader) == 0 || f.Uploader[add["uploader"]]) && allow
}