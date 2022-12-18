package utils

import (
	uuid "github.com/satori/go.uuid"
	"math/big"
	"net/url"
	"regexp"
	"sort"
	"strings"
)

func CreateID() string {
	uid := uuid.NewV4().String()
	var i big.Int
	i.SetString(strings.Replace(uid, "-", "", 4), 16)
	return i.String()
}

func Strip(text string) string {
	text = strings.Replace(text, "\n", " ", -1)
	text = strings.Replace(text, "\t", " ", -1)
	text = strings.Replace(text, "\r", " ", -1)
	text = strings.Replace(text, "\\n", " ", -1)
	re, _ := regexp.Compile("\\\\t")
	text = re.ReplaceAllString(text, "")
	text = strings.TrimSpace(text)
	return text
}

func InString(target string, raw []string) bool {
	sort.Strings(raw)
	index := sort.SearchStrings(raw, target)
	if index < len(raw) && raw[index] == target {
		return true
	}
	return false
}

func JoinString(text []string) string {
	var build strings.Builder
	for _, t := range text {
		build.WriteString(t)
		build.WriteString(" ")
	}
	return build.String()
}

func JoinLink(link, relatives string) (string, error) {
	u, err := url.Parse(relatives)
	if err != nil {
		return "", err
	}
	base, err := url.Parse(link)
	if err != nil {
		return "", err
	}
	return base.ResolveReference(u).String(), nil
}
