package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Service interface {
	GetIpInfo(ip string)(InfoIP, error)
}
type service struct {
	ipApiKey string
}

func NewService(ipApiKey string) (Service, error) {
	return &service{
		ipApiKey: ipApiKey,
	}, nil
}

func (s *service) GetIpInfo(ip string)(InfoIP, error){
	resp, err := http.Get("http://api.ipstack.com/" + ip + "?access_key="+s.ipApiKey)
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var infoip InfoIP
	err = json.Unmarshal(body, &infoip)
	if err != nil {
		log.Println(err)
	}

	return infoip, err
}