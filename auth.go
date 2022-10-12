package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/bettercap/gatt"
	"os"
	"strconv"
	"strings"
)

func auth(p gatt.Peripheral, bytes []byte, c *gatt.Characteristic) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("PIN: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	pin, err := strconv.ParseInt(text, 10, 16)
	if err != nil {
		panic(fmt.Errorf("PIN entry err %s", err))
	}

	pinBase := []byte("\x35\x00\x01\x00")
	pinbuf := make([]byte, 2)
	binary.LittleEndian.PutUint16(pinbuf, uint16(pin))
	pinPkt := append(pinBase, pinbuf...)
	bytes = make([]byte, 20)
	copy(bytes, pinPkt)
	fmt.Printf("Sending PIN: %q\n", bytes)

	err = p.WriteCharacteristic(c, bytes, true)
	if err != nil {
		fmt.Printf("Failed to WriteCharacteristic, err: %s\n", err)
	}
	fmt.Printf("Wrote PIN\n")
	return err
}
