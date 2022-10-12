package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
)

var packetTypes = map[int]func(r *sBProtocolResult){
	SBC_PACKET_HOME_MODE:          parseInt,
	SBC_PACKET_HOME_HSRCSTATE:     parseInt,
	SBC_PACKET_ANODE_VOLTAGE:      parseAnodeVoltage,
	SBC_PACKET_HOME_SENSOR1:       parseDecimal,
	SBC_PACKET_HOME_SENSOR2:       parseDecimal,
	SBC_PACKET_HOME_TEMPERATURE:   parseInt,
	SBC_PACKET_HOME_TEMPNIGHT:     parseInt,
	SBC_PACKET_HOME_TEMPNIGHTCURR: parseInt,
	SBC_PACKET_HOME_TEMPNIGHTLOW:  parseInt,
	SBC_PACKET_HOME_TIME:          parseTime,
	SBC_PACKET_HOME_CAPACITY:      parseInt,
	SBC_PACKET_HOME_ERROR:         parseInt,
	SBC_PACKET_HOME_BOILERMODEL:   parseString,
	SBC_PACKET_HOME_FWVERSION:     parseString,
	SBC_PACKET_HOME_BOILERNAME:    parseString,
	SBC_PACKET_ANODE_PARAMS:       parseAnodeParams,
	SBC_PACKET_HOME_ANTILEGIO:     parseInt,
}

func parseDecimal(r *sBProtocolResult) {
	if t, err := decimal.NewFromString(r.stringData); err == nil {
		r.result = fmt.Sprint(t)
	}
}

func parseAnodeParams(r *sBProtocolResult) {
	if len(r.stringData) < 13 {
		return
	}
	enabled, err := strconv.Atoi(r.stringData[0:3])
	if err != nil {
		return
	}
	errorThreshold, err := strconv.Atoi(r.stringData[3:8])
	if err != nil {
		return
	}
	okThreshold, err := strconv.Atoi(r.stringData[8:13])
	if err != nil {
		return
	}
	r.result = fmt.Sprintf("%d:%d:%d", enabled, errorThreshold, okThreshold)
}

func parseString(r *sBProtocolResult) {
	r.result = r.stringData
}

func parseAnodeVoltage(r *sBProtocolResult) {
	if i, err := decimal.NewFromString(r.stringData); err == nil {
		r.result = fmt.Sprint(i.Div(decimal.NewFromInt(1000)))
	}
}

func parseTime(r *sBProtocolResult) {
	if len(r.stringData) < 10 {
		return
	}
	r.result = r.stringData
}

func parseInt(r *sBProtocolResult) {
	if i, err := strconv.Atoi(r.stringData); err == nil {
		r.result = strconv.Itoa(i)
	}
}
