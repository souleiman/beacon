package sentinel

type filter struct {
    Internal    int         `json:"internal_only"`
    MaxSeeders  int         `json:"maximum_seeders"`
    Resolutions []string    `json:"allowed_resolution"`
    Mediums     []string    `json:"allowed_medium"`
    Codecs      []string    `json:"allowed_codec"`
    Category    []string    `json:"allowed_category"`
    _ResSet     set         `json:"-"`
    _MedSet     set         `json:"-"`
    _CodecSet   set         `json:"-"`
    _CatSet     set         `json:"-"`
}

func (f *filter) PopulateSets() {
    f._ResSet = make(set)
    f._MedSet = make(set)
    f._CodecSet = make(set)
    f._CatSet = make(set)

        for _, item := range f.Resolutions {
            f._ResSet[item] = true
        }

    if f.Mediums != nil {
        for _, item := range f.Mediums {
            f._MedSet[item] = true
        }
    }

    if f.Codecs != nil {
        for _, item := range f.Codecs {
            f._CodecSet[item] = true
        }
    }

    if f.Category != nil {
        for _, item := range f.Category {
            f._CatSet[item] = true
        }
    }
}

func (f *filter) Valid(tor torrent) bool {
    if f.Internal == int(tor.TypeOrigin) {
        return true
    }

    if f.MaxSeeders > 1 && int(tor.Seeders) >= f.MaxSeeders {
        return false
    }

    add := tor.Additional
    return f._ResSet[add["resolution"]] && f._MedSet[add["medium"]] && f._CodecSet[add["codec"]] && f._CatSet[add["category"]]
}