package dns

import (
	"fmt"
	"strings"
)

type RRSet struct {
	Zone  string   `json:"zone" yaml:"zone"`
	RRset []string `json:"rrset" yaml:"rrset"`
}

func ReqRegistUser(config DNSConfig, userId, pubkey string) RRSet {
	return RRSet{Zone: config.UserZone, RRset: []string{fmt.Sprintf("%s. 3600 IN CERT 2 77 2 %s", userId, pubkey)}}
}

func ReqRegistPod(config DNSConfig, userId, podURI string) RRSet {
	return RRSet{Zone: config.UserZone, RRset: []string{fmt.Sprintf("%s. 3600 IN URI 10 1 \"%s\"", userId, podURI)}}
}

func ReqRegistFileURI(config DNSConfig, fileId, userId, fileURI string) RRSet {
	if strings.Contains(userId, config.UserZone) {
		userId = strings.ReplaceAll(userId, config.UserZone, config.PODZone)
	}

	if strings.Contains(fileId, config.PODZone) {
		fileId = strings.ReplaceAll(fileId, config.PODZone, config.UserZone)
	}
	return RRSet{Zone: fmt.Sprintf("%s", userId), RRset: []string{fmt.Sprintf("%s. 3600 IN URI 10 1  \"%s\"", fileId, fileURI)}}
}

func ReqRegistFile(config DNSConfig, fileId, userId, hash, hash_func string) RRSet {
	oldUserId := userId
	if strings.Contains(userId, config.UserZone) {
		userId = strings.ReplaceAll(userId, config.UserZone, config.PODZone)
	}

	if strings.Contains(fileId, config.PODZone) {
		fileId = strings.ReplaceAll(fileId, config.PODZone, config.UserZone)
	}

	oldUserId = strings.ReplaceAll(oldUserId, "."+config.UserZone, "")

	return RRSet{Zone: fmt.Sprintf("%s", userId), RRset: []string{fmt.Sprintf("%s. 3600 IN TXT \"user=%s,algo=%s,hash=%s\"", fileId, oldUserId, hash_func, hash)}}
}

type Zone struct {
	Name string   `json:"name" yaml:"name"`
	IPs  []string `json:"ips" yaml:"ips"`
}

func ReqZone(userID string, cfg DNSConfig) Zone {
	if strings.Contains(userID, cfg.UserZone) {
		userID = strings.ReplaceAll(userID, cfg.UserZone, cfg.PODZone)
	}
	return Zone{Name: fmt.Sprintf("%s", userID), IPs: []string{cfg.IPAB}}
}

type Forward struct {
	Name string `json:"name" yaml:"name"`
	Addr string `json:"addr" yaml:"addr"`
}

func ReqForward(userID string, cfg DNSConfig) Forward {
	if strings.Contains(userID, cfg.UserZone) {
		userID = strings.ReplaceAll(userID, cfg.UserZone, cfg.PODZone)
	}
	return Forward{Name: fmt.Sprintf("%s", userID), Addr: fmt.Sprintf("%s:5555", cfg.IPAB)}
}
