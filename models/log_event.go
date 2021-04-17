package models

import (
	"fmt"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FloatString float64

func (a FloatString) UnmarshalJSON(b []byte) error {
	in := string(b)
	if strings.ToLower(in) == "unknown" {
		a = 0.0
		return nil
	}

	//var s string
	//if err := json.Unmarshal(b, &s); err != nil {
	//	return fmt.Errorf("json unmarshall: %w", err)
	//}

	aa, err := strconv.ParseFloat(in, 64)
	if err != nil {
		return fmt.Errorf("parse float: %w", err)
	}
	a = FloatString(aa)
	return nil
}

type LogEvent struct {
	mgoID   primitive.ObjectID `json:"mgoID,omitempty" bson:"_id,omitempty"`
	AppEui  string             `json:"app_eui"`
	Decoded struct {
		Payload struct {
			Accuracy  float64 `json:"Accuracy"`
			Altitude  float64 `json:"Altitude"`
			Latitude  float64 `json:"Latitude"`
			Longitude float64 `json:"Longitude"`
			Pitch     float64 `json:"Pitch"`
			Roll      float64 `json:"Roll"`
		} `json:"payload"`
		Status string `json:"status"`
	} `json:"decoded"`
	DevEui      string `json:"dev_eui"`
	Devaddr     string `json:"devaddr"`
	DownlinkURL string `json:"downlink_url"`
	Fcnt        int    `json:"fcnt"`
	Hotspots    []struct {
		Channel    float64     `json:"channel"`
		Frequency  float64     `json:"frequency"`
		ID         string      `json:"id"`
		Lat        FloatString `json:"lat"`
		Long       FloatString `json:"long"`
		Name       string      `json:"name"`
		ReportedAt int64       `json:"reported_at"`
		Rssi       float64     `json:"rssi"`
		Snr        float64     `json:"snr"`
		Spreading  string      `json:"spreading"`
		Status     string      `json:"status"`
	} `json:"hotspots"`
	ID       int `json:"-"`
	Metadata struct {
		Labels []struct {
			ID             string `json:"id"`
			Name           string `json:"name"`
			OrganizationID string `json:"organization_id"`
		} `json:"labels"`
		MultiBuy       float64 `json:"multi_buy"`
		OrganizationID string  `json:"organization_id"`
	} `json:"metadata"`
	Name        string  `json:"name"`
	Payload     string  `json:"payload"`
	PayloadSize float64 `json:"payload_size"`
	Port        float64 `json:"port"`
	ReportedAt  int64   `json:"reported_at"`
	UUID        string  `json:"uuid"`
}
