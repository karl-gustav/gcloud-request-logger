package main

import "time"

type Lora struct {
	DevEUIUplink struct {
		Time       time.Time `json:"Time"`
		DevEUI     string    `json:"DevEUI"`
		FPort      int       `json:"FPort"`
		FCntUp     int       `json:"FCntUp"`
		ADRbit     int       `json:"ADRbit"`
		MType      int       `json:"MType"`
		FCntDn     int       `json:"FCntDn"`
		PayloadHex string    `json:"payload_hex"`
		MicHex     string    `json:"mic_hex"`
		Lrcid      string    `json:"Lrcid"`
		LrrRSSI    float64   `json:"LrrRSSI"`
		LrrSNR     float64   `json:"LrrSNR"`
		SpFact     int       `json:"SpFact"`
		SubBand    string    `json:"SubBand"`
		Channel    string    `json:"Channel"`
		DevLrrCnt  int       `json:"DevLrrCnt"`
		Lrrid      string    `json:"Lrrid"`
		Late       int       `json:"Late"`
		LrrLAT     float64   `json:"LrrLAT"`
		LrrLON     float64   `json:"LrrLON"`
		Lrrs       struct {
			Lrr []struct {
				Lrrid   string  `json:"Lrrid"`
				Chain   int     `json:"Chain"`
				LrrRSSI float64 `json:"LrrRSSI"`
				LrrSNR  float64 `json:"LrrSNR"`
				LrrESP  float64 `json:"LrrESP"`
			} `json:"Lrr"`
		} `json:"Lrrs"`
		CustomerID   string `json:"CustomerID"`
		CustomerData struct {
			Alr struct {
				Pro string `json:"pro"`
				Ver string `json:"ver"`
			} `json:"alr"`
		} `json:"CustomerData"`
		ModelCfg   string  `json:"ModelCfg"`
		InstantPER float64 `json:"InstantPER"`
		MeanPER    float64 `json:"MeanPER"`
		DevAddr    string  `json:"DevAddr"`
		TxPower    float64 `json:"TxPower"`
		NbTrans    int     `json:"NbTrans"`
	} `json:"DevEUI_uplink"`
}
