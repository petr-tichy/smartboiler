package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
	"log"
	"strconv"
)

func handleNotify(b []byte) (r sBProtocolResult) {
	var err error
	if len(b) < 2 {
		return
	}
	r.rqType, err = strconv.Atoi(string(b[:2]))
	if err != nil {
		return
	}
	switch r.rqType {
	case SBC_PACKET_NIGHT_GETDAYS:
		if len(b) < 14 {
			return
		}
		r.byteData = b[2:14]
	case SBC_PACKET_NIGHT_GETDAYS2:
		if len(b) < 11 {
			return
		}
		r.byteData = b[2:11]
	case SBC_PACKET_GLOBAL_CONFIRMUID:
		if len(b) < 20 {
			return
		}
		r.uid = b[2]
		r.byteData = b[4:20]
	case SBC_PACKET_HOLIDAY_GET:
		if !decodeIntArr(b, r) {
			return
		}
	case SBC_PACKET_GLOBAL_FIRSTLOG:
		if !decodeIntArr(b, r) {
			return
		}
	case SBC_PACKET_GLOBAL_NEXTLOG:
		if !decodeIntArr(b, r) {
			return
		}
	}

	b = b[2:]
	b = b[:bytes.IndexByte(b, 0)]
	e, err := ianaindex.MIME.Encoding("US-ASCII")
	if err != nil {
		log.Println(err)
	}
	str, _, err := transform.String(e.NewDecoder(), string(b))
	if err != nil || len(str) == 0 {
		return
	}
	r.stringData = str

	if name, ok := topics[r.rqType]; ok {
		r.rqName = name
	} else {
		log.Println("unknown rqType:", r.rqType)
	}

	if f, ok := packetTypes[r.rqType]; ok {
		f(&r)
	} else {
		log.Println("unknown parser", r.rqName)
	}

	fmt.Printf("received: %s: %d | %d | % X | %q | %q | %q\n", r.rqName, r.rqType, r.uid, r.byteData, r.byteData, r.stringData, r.result)
	return r
}

func decodeIntArr(b []byte, r sBProtocolResult) bool {
	if len(b) < 18 {
		return false
	}
	i0, err := strconv.Atoi(string(b[2:8]))
	if err != nil {
		return false
	}
	i1, err := strconv.Atoi(string(b[10:18]))
	if err != nil {
		return false
	}
	r.intArrData[0] = i0
	r.intArrData[1] = i1
	return true
}
