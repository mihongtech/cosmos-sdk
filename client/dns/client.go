package dns

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/tendermint/tendermint/libs/log"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	AddRRSet   = "/AddRRset"
	AddZone    = "/AddZone"
	AddForward = "/AddForward"
)

type Client struct {
	log log.Logger
	cfg DNSConfig
}

func NewClient(logger log.Logger, cfg DNSConfig) *Client {
	return &Client{log: logger, cfg: cfg}
}

func (c Client) post(url string, reqBody io.Reader) (result []byte, err error) {
	httpReq, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		c.log.Error("NewRequest fail", "url:", url, "reqBody: ", reqBody, "err: ", err)
		return nil, err
	}
	httpReq.Header.Add("Content-Type", "application/json")

	// DO: HTTP请求
	httpRsp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		c.log.Error("do http fail", "url:", url, "reqBody: ", reqBody, "err: ", err)
		return nil, err
	}
	defer httpRsp.Body.Close()

	// Read: HTTP结果
	result, err = ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		c.log.Error("ReadAll failed", "url:", url, "reqBody: ", reqBody, "err: ", err)
		return nil, err
	}
	return
}

func (c Client) addRRSet(url string, set RRSet) error {
	c.log.Info("addRRSet", "url", url, "request", set)
	buff, err := json.Marshal(set)
	if err != nil {
		c.log.Error("Marshal RequestParam failed", "err: ", err)
		return err
	}

	reqBody := strings.NewReader(string(buff))
	result, err := c.post(url, reqBody)
	if err != nil {
		c.log.Error("Add RRSet Post failed", "err: ", err)
		return err
	}

	if bytes.Compare(buff, result) != 0 {
		c.log.Error("addRRSet", "result", string(result))
		return errors.New("failed add rrset")
	}
	c.log.Info("addRRSet", "result", string(result))
	return nil
	//var rr RRSet
	//if err := json.Unmarshal(result, &rr); err != nil {
	//	return err
	//}
}

func (c Client) addZone(url string, set Zone) error {
	c.log.Info("addZone", "url", url, "request", set)
	buff, err := json.Marshal(set)
	if err != nil {
		c.log.Error("Marshal RequestParam fail, err:%v", err)
		return err
	}

	reqBody := strings.NewReader(string(buff))
	result, err := c.post(url, reqBody)
	if err != nil {
		c.log.Error("Add ZONE Post failed, err:%v", err)
		return err
	}
	if bytes.Compare(buff, result) != 0 {
		c.log.Error("addZone", "result", string(result))
		return errors.New("failed add rrset")
	}
	c.log.Info("addZone", "result", string(result))
	return nil
}

func (c Client) addForward(url string, set Forward) error {
	c.log.Info("addForward", "url", url, "request", set)
	buff, err := json.Marshal(set)
	if err != nil {
		c.log.Error("Marshal RequestParam fail, err:%v", err)
		return err
	}

	reqBody := strings.NewReader(string(buff))
	result, err := c.post(url, reqBody)
	if err != nil {
		c.log.Error("Add Forward Post failed, err:%v", err)
		return err
	}
	//if bytes.Compare(buff, result) != 0 {
	//	c.log.Error("addForward", "result", string(result))
	//	return errors.New("failed add rrset")
	//}
	c.log.Info("addForward", "result", string(result))
	return nil
}

func (c Client) RegisterUser(userId, pubKey string) error {
	urlAA := c.cfg.ServerAA + AddRRSet
	RRSet := ReqRegistUser(c.cfg, userId, pubKey)
	return c.addRRSet(urlAA, RRSet)
}

func (c Client) RegisterPOD(userId, podURI string) error {
	urlAA := c.cfg.ServerAA + AddRRSet
	RRSet := ReqRegistPod(c.cfg, userId, podURI)
	if err := c.addRRSet(urlAA, RRSet); err != nil {
		return err
	}

	urlAB := c.cfg.ServerAB + AddZone
	zone := ReqZone(userId, c.cfg)
	if err := c.addZone(urlAB, zone); err != nil {
		return err
	}

	urlCC := c.cfg.ServerCC + AddForward
	forward := ReqForward(userId, c.cfg)
	if err := c.addForward(urlCC, forward); err != nil {
		return err
	}
	return nil
}

func (c Client) Upload(fileId, userId, podURI string, hash, hash_func string) error {
	urlAB := c.cfg.ServerAB + AddRRSet
	RRSet := ReqRegistFileURI(c.cfg, fileId, userId, podURI)
	if err := c.addRRSet(urlAB, RRSet); err != nil {
		return err
	}

	RRSet = ReqRegistFile(c.cfg, fileId, userId, hash, hash_func)
	if err := c.addRRSet(urlAB, RRSet); err != nil {
		return err
	}

	return nil
}
