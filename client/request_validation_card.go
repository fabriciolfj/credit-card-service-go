package client

import (
	json2 "encoding/json"
	"errors"
	"fmt"
	"github.com/fabriciolfj/credit-card-service-go/model"
	"github.com/magiconair/properties"
	"net/http"
)

type RequestCard struct {
	client http.Client
	url    string
}

func ProvideRequestCard() *RequestCard {
	config, err := properties.LoadFile("config.properties", properties.UTF8)

	if err != nil {
		panic("properties not found")
	}

	return &RequestCard{
		client: http.Client{},
		url:    config.GetString("client.approve.card", ""),
	}
}

func (r *RequestCard) FindApprove(code string) (*model.CardCustomerApproveDto, error) {
	resp, err := http.Get(r.url + "/" + code)

	if err != nil {
		panic(errors.New("http get error, details " + err.Error()))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad request, code %d", resp.StatusCode)
	}

	var card model.CardCustomerApproveDto
	if err := json2.NewDecoder(resp.Body).Decode(&card); err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &card, nil
}
