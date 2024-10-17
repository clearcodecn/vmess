package vmess

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

type Link struct {
	Version int    `json:"v"`    // version
	Name    string `json:"ps"`   // name
	Address string `json:"add"`  // server address, ip or domain
	Port    string `json:"port"` // server port.
	Id      string `json:"id"`   // server id
	Aid     string `json:"aid"`  // alert ID
	Network string `json:"net"`  // protocol: tcp/kcp/ws/h2/quic
	Type    string `json:"type"` // none\http\srtp\utp\wechat-video
	Host    string `json:"host"` // host (when http, ws, h2) || quic
	Path    string `json:"path"` // path(ws, h2)
	Tls     string `json:"tls"`  // tls settings.
}

func (l *Link) ToVmessLink() string {
	data, _ := json.Marshal(l)
	return "vmess://" + base64.StdEncoding.EncodeToString(data)
}

func FromVmessLink(link string) (*Link, error) {
	b64 := strings.TrimPrefix(link, "vmess://")
	b64Decoded, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, err
	}
	var vmess Link
	err = json.Unmarshal(b64Decoded, &vmess)
	if err != nil {
		return nil, err
	}
	return &vmess, nil
}
