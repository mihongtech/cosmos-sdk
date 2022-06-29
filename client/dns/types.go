package dns

import "fmt"

type RRSet struct {
	Zone  string   `json:"zone" yaml:"zone"`
	RRset []string `json:"rrset" yaml:"rrset"`
}

func ReqRegistUser(config DNSConfig, userId, pubkey string) RRSet {
	return RRSet{Zone: config.UserZone, RRset: []string{fmt.Sprintf("%s.%s. 3600 IN CERT 2 77 2 %s", userId, config.UserZone, pubkey)}}
}

func ReqRegistPod(config DNSConfig, userId, podURI string) RRSet {
	return RRSet{Zone: config.UserZone, RRset: []string{fmt.Sprintf("%s.%s. 3600 IN URI 10 1 \"%s\"", userId, config.UserZone, podURI)}}
}

func ReqRegistFileURI(config DNSConfig, fileId, userId, fileURI string) RRSet {
	return RRSet{Zone: fmt.Sprintf("%s.%s", userId, config.PODZone), RRset: []string{fmt.Sprintf("%s.%s.%s. 3600 IN URI 10 1  \"%s\"", fileId, userId, config.PODZone, fileURI)}}
}

func ReqRegistFile(config DNSConfig, fileId, userId, hash, hash_func string) RRSet {
	return RRSet{Zone: fmt.Sprintf("%s.%s", userId, config.PODZone), RRset: []string{fmt.Sprintf("%s.%s.%s. 3600 IN TXT \"user=%s,algo=%s,hash=%s\"", fileId, userId, config.PODZone, userId, hash_func, hash)}}
}

type Zone struct {
	Name string   `json:"name" yaml:"name"`
	IPs  []string `json:"ips" yaml:"ips"`
}

func ReqZone(userID string, cfg DNSConfig) Zone {
	return Zone{Name: fmt.Sprintf("%s.%s", userID, cfg.PODZone), IPs: []string{cfg.IPAB}}
}

type Forward struct {
	Name string `json:"name" yaml:"name"`
	Addr string `json:"addr" yaml:"addr"`
}

func ReqForward(userID string, cfg DNSConfig) Forward {
	return Forward{Name: fmt.Sprintf("%s.%s", userID, cfg.PODZone), Addr: fmt.Sprintf("%s:5555", cfg.IPAB)}
}
