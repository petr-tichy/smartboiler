package main

import "github.com/bettercap/gatt"

const (
	SBC_PACKET_NONE                          = 0
	SBC_PACKET_LOG_MSG                       = 1
	SBC_PACKET_HOME_BOILERMODEL              = 2
	SBC_PACKET_HOME_FWVERSION                = 3
	SBC_PACKET_HOME_MODE                     = 4
	SBC_PACKET_HOME_ERROR                    = 5
	SBC_PACKET_HOME_HSRCSTATE                = 6
	SBC_PACKET_HOME_SENSOR1                  = 7
	SBC_PACKET_HOME_SENSOR2                  = 8
	SBC_PACKET_HOME_TEMPERATURE              = 9
	SBC_PACKET_HOME_TEMPNIGHT                = 10
	SBC_PACKET_HOME_TIME                     = 11
	SBC_PACKET_HOME_ALL                      = 12
	SBC_PACKET_HOME_SETTIME                  = 13
	SBC_PACKET_HOME_SETNORMALTEMPERATURE     = 14
	SBC_PACKET_HOME_TEMPNIGHTLOW             = 15
	SBC_PACKET_HOME_SETTEMPNIGHT             = 16
	SBC_PACKET_HOME_SETTEMPNIGHTLOW          = 17
	SBC_PACKET_HOME_SETMODE                  = 18
	SBC_PACKET_HOME_SETNORMALMODE            = 19
	SBC_PACKET_HDO_ONOFF                     = 20
	SBC_PACKET_HDO_SELECTION_A               = 21
	SBC_PACKET_HDO_SELECTION_B               = 22
	SBC_PACKET_HDO_SELECTION_DP              = 23
	SBC_PACKET_HDO_FREQUENCY                 = 24
	SBC_PACKET_HDO_SETTING                   = 25
	SBC_PACKET_HDO_ALL                       = 26
	SBC_PACKET_HDO_SET_ONOFF                 = 27
	SBC_PACKET_HDO_SET_SELECTION_A           = 28
	SBC_PACKET_HDO_SET_SELECTION_B           = 29
	SBC_PACKET_HDO_SET_SELECTION_DP          = 30
	SBC_PACKET_HDO_SET_FREQUENCY             = 31
	SBC_PACKET_HDO_LASTHDOTIME               = 32
	SBC_PACKET_HDO_LESSEXPTARIFFAVAILABLENOW = 33
	SBC_PACKET_HDO_INFO                      = 34
	SBC_PACKET_HDO_MANUAL_SET                = 35
	SBC_PACKET_HDO_MANUAL_GET                = 36
	SBC_PACKET_NIGHT_GETDAY                  = 40
	SBC_PACKET_NIGHT_SAVEDAY                 = 41
	SBC_PACKET_NIGHT_SAVEMINMAX              = 42
	SBC_PACKET_NIGHT_SAVEDAYS                = 43
	SBC_PACKET_NIGHT_GETDAYS                 = 44
	SBC_PACKET_NIGHT_SAVEDAYS2               = 45
	SBC_PACKET_NIGHT_GETDAYS2                = 46
	SBC_PACKET_GLOBAL_STARTSIMULATION        = 50
	SBC_PACKET_GLOBAL_CONFIRMUID             = 51
	SBC_PACKET_GLOBAL_ERRORUID               = 52
	SBC_PACKET_GLOBAL_PAIRPIN                = 53
	SBC_PACKET_GLOBAL_FIRSTLOG               = 54
	SBC_PACKET_GLOBAL_NEXTLOG                = 55
	SBC_PACKET_GLOBAL_RESETBERR              = 56
	SBC_PACKET_GLOBAL_PINRESULT              = 57
	SBC_PACKET_GLOBAL_DEVICEBONDED           = 58
	SBC_PACKET_HOME_FWRESET                  = 59
	SBC_PACKET_HOLIDAY_GET                   = 60
	SBC_PACKET_HOLIDAY_SET                   = 61
	SBC_PACKET_HOLIDAY_ENABLE                = 62
	SBC_PACKET_HOLIDAY_DISABLE               = 63
	SBC_PACKET_HOLIDAY_DELETE                = 64
	SBC_PACKET_HOLIDAY_ENABLED               = 65
	SBC_PACKET_RQ_GLOBAL_MAC                 = 68
	SBC_PACKET_RQ_UI_BUTTONUP                = 70
	SBC_PACKET_RQ_UI_BUTTONDOWN              = 71
	SBC_PACKET_RQ_UI_D7SEG                   = 72
	SBC_PACKET_POWERCONS_RESET               = 73
	SBC_PACKET_POWERCONS_OBTAIN              = 74
	SBC_PACKET_HOME_ANTILEGIO                = 76
	SBC_PACKET_HOME_NEWERR                   = 77
	SBC_PACKET_HOME_BOILERNAME               = 80
	SBC_PACKET_HOME_SETBOILERNAME            = 81
	SBC_PACKET_HOME_TEMPNIGHTCURR            = 82
	SBC_PACKET_HOME_CAPACITY                 = 83
	SBC_PACKET_HOME_SETCAPACITY              = 84
	SBC_PACKET_HOME_FWBEGIN                  = 86
	SBC_PACKET_HOME_FWCONFIRM                = 87
	SBC_PACKET_HOME_FWCHECK                  = 88
	SBC_PACKET_HOME_FWCOPY                   = 89
	SBC_PACKET_STATISTICS_WEEK               = 90
	SBC_PACKET_STATISTICS_YEAR               = 91
	SBC_PACKET_STATISTICS_RESET              = 92
	SBC_PACKET_STATISTICS_GETALL             = 93
	SBC_PACKET_ANODE_VOLTAGE                 = 94
	SBC_PACKET_ANODE_PARAMS                  = 95
)

const (
	SBCMODE_STOP      = 0
	SBCMODE_NORMAL    = 1
	SBCMODE_HDO       = 2
	SBCMODE_SMART     = 3
	SBCMODE_SMARTHDO  = 4
	SBCMODE_ANTIFROST = 5
	SBCMODE_NIGHT     = 6
	SBCMODE_TEST      = 7
	SBCMODE_HOLIDAY   = 8
)

var (
	sBCoreServiceUUID                        = gatt.UUID16(0x1899)
	sBCoreCharacteristicUUID                 = gatt.UUID16(0x2b99)
	logServiceUUID                           = gatt.UUID16(0x1898)
	logCharacteristicUUID                    = gatt.UUID16(0x2b98)
	clientCharacteristicConfigDescriptorUUID = gatt.UUID16(0x2902)
	deviceInformationServiceUUID             = gatt.UUID16(0x180a)

	topics = map[int]string{
		SBC_PACKET_LOG_MSG:                       "log",
		SBC_PACKET_HOME_BOILERMODEL:              "model",
		SBC_PACKET_HOME_FWVERSION:                "fw_version",
		SBC_PACKET_HOME_MODE:                     "mode",
		SBC_PACKET_HOME_ERROR:                    "error",
		SBC_PACKET_HOME_HSRCSTATE:                "heating",
		SBC_PACKET_HOME_SENSOR1:                  "sensor1",
		SBC_PACKET_HOME_SENSOR2:                  "sensor2",
		SBC_PACKET_HOME_TEMPERATURE:              "temperature",
		SBC_PACKET_HOME_TEMPNIGHT:                "temp_night",
		SBC_PACKET_HOME_TIME:                     "time",
		SBC_PACKET_HOME_ALL:                      "all",
		SBC_PACKET_HOME_TEMPNIGHTLOW:             "temp_night_low",
		SBC_PACKET_HDO_ONOFF:                     "hdo_on_off",
		SBC_PACKET_HDO_SELECTION_A:               "hdo_A",
		SBC_PACKET_HDO_SELECTION_B:               "hdo_B",
		SBC_PACKET_HDO_SELECTION_DP:              "hdo_DP",
		SBC_PACKET_HDO_FREQUENCY:                 "hdo_freq",
		SBC_PACKET_HDO_SETTING:                   "hdo_setting",
		SBC_PACKET_HDO_ALL:                       "hdo_all",
		SBC_PACKET_HDO_LASTHDOTIME:               "hdo_last_time",
		SBC_PACKET_HDO_LESSEXPTARIFFAVAILABLENOW: "hdo_now",
		SBC_PACKET_HDO_INFO:                      "hdo_info",
		//SBC_PACKET_HDO_MANUAL_SET                = 35
		//SBC_PACKET_HDO_MANUAL_GET                = 36
		//SBC_PACKET_NIGHT_GETDAY                  = 40
		//SBC_PACKET_NIGHT_SAVEDAY                 = 41
		//SBC_PACKET_NIGHT_SAVEMINMAX              = 42
		//SBC_PACKET_NIGHT_SAVEDAYS                = 43
		//SBC_PACKET_NIGHT_GETDAYS                 = 44
		//SBC_PACKET_NIGHT_SAVEDAYS2               = 45
		//SBC_PACKET_NIGHT_GETDAYS2                = 46
		//SBC_PACKET_GLOBAL_STARTSIMULATION        = 50
		SBC_PACKET_GLOBAL_CONFIRMUID: "confirm_uid",
		SBC_PACKET_GLOBAL_ERRORUID:   "error_uid",
		SBC_PACKET_GLOBAL_PAIRPIN:    "pair_pin",
		//SBC_PACKET_GLOBAL_FIRSTLOG               = 54
		//SBC_PACKET_GLOBAL_NEXTLOG                = 55
		//SBC_PACKET_GLOBAL_RESETBERR              = 56
		//SBC_PACKET_GLOBAL_PINRESULT              = 57
		//SBC_PACKET_GLOBAL_DEVICEBONDED           = 58
		//SBC_PACKET_HOME_FWRESET                  = 59
		//SBC_PACKET_HOLIDAY_GET                   = 60
		//SBC_PACKET_HOLIDAY_SET                   = 61
		//SBC_PACKET_HOLIDAY_ENABLE                = 62
		//SBC_PACKET_HOLIDAY_DISABLE               = 63
		//SBC_PACKET_HOLIDAY_DELETE                = 64
		//SBC_PACKET_HOLIDAY_ENABLED               = 65
		SBC_PACKET_RQ_GLOBAL_MAC: "mac",
		//SBC_PACKET_RQ_UI_BUTTONUP                = 70
		//SBC_PACKET_RQ_UI_BUTTONDOWN              = 71
		//SBC_PACKET_RQ_UI_D7SEG                   = 72
		//SBC_PACKET_POWERCONS_RESET               = 73
		//SBC_PACKET_POWERCONS_OBTAIN              = 74
		SBC_PACKET_HOME_ANTILEGIO:     "antilegionella",
		SBC_PACKET_HOME_NEWERR:        "new_error",
		SBC_PACKET_HOME_BOILERNAME:    "name",
		SBC_PACKET_HOME_TEMPNIGHTCURR: "temp_night_current",
		SBC_PACKET_HOME_CAPACITY:      "capacity",
		//SBC_PACKET_STATISTICS_WEEK               = 90
		//SBC_PACKET_STATISTICS_YEAR               = 91
		//SBC_PACKET_STATISTICS_RESET              = 92
		SBC_PACKET_STATISTICS_GETALL: "statistics_all",
		SBC_PACKET_ANODE_VOLTAGE:     "anode_voltage",
		SBC_PACKET_ANODE_PARAMS:      "anode_params",
	}
)
