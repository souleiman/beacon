package sentinel

import (
    "regexp"
)


func init() {
    re = regexp.MustCompile(regex)
}
func announce_matcher(text string) params {
    matches := re.FindStringSubmatch(text)

    result := make(map[string]string)
    for i, name := range re.SubexpNames() {
        if name == "" {
            continue
        }
        result[name] = matches[i]
    }

    return result
}