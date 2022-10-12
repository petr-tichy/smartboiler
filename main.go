package main

import (
	"encoding/binary"
	"fmt"
	"github.com/bettercap/gatt"
	"github.com/bettercap/gatt/examples/option"
	"log"
	"time"
)

type sBProtocolRequest struct {
	rqType  int
	uid     int
	rawData []byte
}

type sBProtocolResult struct {
	rqType     int
	uid        byte
	rawData    []byte
	stringData string
	byteData   []byte
	intArrData []int
	rqName     string
	result     string
}

var in = make(chan sBProtocolRequest)

func onStateChanged(d gatt.Device, s gatt.State) {
	fmt.Println("State:", s)
	switch s {
	case gatt.StatePoweredOn:
		startScan(d)
		return
	default:
		d.StopScanning()
	}
}

func startScan(d gatt.Device) {
	fmt.Println("Scanning...")
	d.Scan([]gatt.UUID{logCharacteristicUUID, deviceInformationServiceUUID}, true)
}

func onPeripheralDiscovered(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {
	if !(len(a.Services) == 2 &&
		gatt.UUIDContains(a.Services, logCharacteristicUUID) &&
		gatt.UUIDContains(a.Services, deviceInformationServiceUUID)) {
		return
	}

	p.Device().StopScanning()
	p.Device().Connect(p)
}

func onPeripheralConnected(p gatt.Peripheral, err error) {
	fmt.Println("Connected")
	defer p.Device().CancelConnection(p)

	if err := p.SetMTU(200); err != nil {
		fmt.Printf("Failed to set MTU, err: %s\n", err)
	}

	// Discovery services

	ss, err := p.DiscoverServices([]gatt.UUID{sBCoreServiceUUID, logServiceUUID})
	if err != nil {
		fmt.Printf("Failed to discover services, err: %s\n", err)
		return
	}

	for _, s := range ss {
		// Discovery characteristics
		cs, err := p.DiscoverCharacteristics([]gatt.UUID{sBCoreCharacteristicUUID, logCharacteristicUUID}, s)
		if err != nil {
			fmt.Printf("Failed to discover characteristics, err: %s\n", err)
			continue
		}

		for _, c := range cs {
			_, err := p.DiscoverDescriptors([]gatt.UUID{clientCharacteristicConfigDescriptorUUID}, c)
			if err != nil {
				fmt.Printf("Failed to discover descriptors, err: %s\n", err)
				continue
			}

			if c.UUID().Equal(logCharacteristicUUID) && c.Properties()&gatt.CharNotify != 0 {
				f := func(c *gatt.Characteristic, b []byte, err error) {
					handleNotify(b)
				}
				if err := p.SetNotifyValue(c, f); err != nil {
					fmt.Printf("Failed to subscribe characteristic, err: %s\n", err)
					continue
				}
			}

			if c.UUID().Equal(sBCoreCharacteristicUUID) && c.Properties()&gatt.CharWrite != 0 {
				send(p, c, sBProtocolRequest{
					rqType:  0x44,
					uid:     0x44,
					rawData: []byte("\xde\xad\xbe\xef\x12\x34"),
				})

				// err := auth(p, bytes, c)

				go checkSmart(in)

				for msg := range in {
					send(p, c, msg)
				}
			}
		}
		fmt.Println()
	}

}

func send(p gatt.Peripheral, c *gatt.Characteristic, msg sBProtocolRequest) {
	pkt := make([]byte, 2)
	binary.LittleEndian.PutUint16(pkt, uint16(msg.rqType))
	t := make([]byte, 2)
	binary.LittleEndian.PutUint16(t, uint16(msg.uid))
	pkt = append(pkt, t...)
	pkt = append(pkt, msg.rawData...)
	err := p.WriteCharacteristic(c, pkt, true)
	if err != nil {
		fmt.Printf("Failed to WriteCharacteristic, err: %s\n", err)
	}
	rqName := "unknown"
	if _rqName, ok := topics[msg.rqType]; ok {
		rqName = _rqName
	}
	fmt.Printf("Wrote msg: %s | % X\n", rqName, pkt)
	time.Sleep(50 * time.Millisecond)
}

func onPeripheralDisconnected(p gatt.Peripheral, err error) {
	close(in)
	startScan(p.Device())
}

func checkSmart(ch chan sBProtocolRequest) {
	for {
		ch <- sBProtocolRequest{SBC_PACKET_HOME_ERROR, SBC_PACKET_HOME_ERROR, nil}
		ch <- sBProtocolRequest{SBC_PACKET_HOME_NEWERR, SBC_PACKET_HOME_NEWERR, nil}
		ch <- sBProtocolRequest{SBC_PACKET_HOME_HSRCSTATE, SBC_PACKET_HOME_HSRCSTATE, nil}
		ch <- sBProtocolRequest{SBC_PACKET_HOME_SENSOR1, SBC_PACKET_HOME_SENSOR1, nil}
		ch <- sBProtocolRequest{SBC_PACKET_HOME_SENSOR2, SBC_PACKET_HOME_SENSOR2, nil}
		ch <- sBProtocolRequest{SBC_PACKET_STATISTICS_GETALL, SBC_PACKET_STATISTICS_GETALL, nil}
		ch <- sBProtocolRequest{SBC_PACKET_ANODE_VOLTAGE, SBC_PACKET_ANODE_VOLTAGE, nil}
		ch <- sBProtocolRequest{SBC_PACKET_HOME_ANTILEGIO, SBC_PACKET_HOME_ANTILEGIO, nil}
		ch <- sBProtocolRequest{SBC_PACKET_HOME_ALL, SBC_PACKET_HOME_ALL, nil}
		time.Sleep(5 * time.Second)
	}
}

func main() {
	d, err := gatt.NewDevice(option.DefaultClientOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s\n", err)
		return
	}

	// Register handlers.
	d.Handle(
		gatt.PeripheralDiscovered(onPeripheralDiscovered),
		gatt.PeripheralConnected(onPeripheralConnected),
		gatt.PeripheralDisconnected(onPeripheralDisconnected),
	)

	if err := d.Init(onStateChanged); err != nil {
		log.Fatalf("Failed init", err)
	}

	select {}
}
