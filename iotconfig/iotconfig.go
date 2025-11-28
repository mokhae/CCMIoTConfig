package iotconfig

import (
	"encoding/json"
	"os"
)

type Comm_Mode int

const (
	COMM_MODE_ETH Comm_Mode = iota
	COMM_MOMDE_SERIAL
)

func (s Comm_Mode) String() string {
	switch s {
	case COMM_MODE_ETH:
		return "Ethernet"
	case COMM_MOMDE_SERIAL:
		return "Serial"
	default:
		return "Unknown"
	}
}

type Slave_Mode int

const (
	USE_MASTER Slave_Mode = iota
	USE_SLAVE
)

func (s Slave_Mode) String() string {
	switch s {
	case USE_MASTER:
		return "Use_Master"
	case USE_SLAVE:
		return "Use_Slave"
	default:
		return "Unknown"
	}
}

type Modbus_Serial_Type int

const (
	MB_SERIAL_RTU Modbus_Serial_Type = iota
	MB_SERIAL_ASCII
	MB_SERIAL_RTU_OVER_TCP
	MB_SERIAL_ASCII_OVER_TCP
)

func (s Modbus_Serial_Type) String() string {
	switch s {
	case MB_SERIAL_RTU:
		return "MODBUS RTU"
	case MB_SERIAL_ASCII:
		return "MODBUS ASCII"
	case MB_SERIAL_RTU_OVER_TCP:
		return "MODBUS RTU OVER TCP"
	case MB_SERIAL_ASCII_OVER_TCP:
		return "MODBUS ASCII OVER TCP"
	default:
		return "Unknown"
	}
}

type MC3E_CONFIG struct {
	DEVICES  []MC3E_DEVICE `json:"DEVICES"`
	USE_COMM bool          `json:"USE_COMM"`
}

type MODBUS_CONFIG struct {
	DEVICES  []MODBUS_DEVICE `json:"DEVICES"`
	USE_COMM bool            `json:"USE_COMM"`
}

type MODBUS_DEVICE struct {
	DEVICE_ID          string             `json:"DEVICE_ID"`
	COMM_MODE          Comm_Mode          `json:"COMM_MODE"`
	USE_SLAVE          Slave_Mode         `json:"USE_SLAVE"`
	SLAVE_ID           byte               `json:"SLAVE_ID"`
	MODBUS_SERIAL_TYPE Modbus_Serial_Type `json:"MODBUS_TYPE"`
	ETHERNET           ETH_ST             `json:"ETHERNET"`
	SERIAL             SERIAL_ST          `json:"SERIAL"`
	TIMEOUT            TIMEOUT_ST         `json:"TIMEOUT"`
	CHANNELS           []MB_CHANNEL       `json:"CHANNELS"`
	ENABLED            bool               `json:"ENABLED"`
	USE_DEVICE         bool               `json:"USE_DEVICE"`
}

type MC3E_DEVICE struct {
	DEVICE_ID  string         `json:"DEVICE_ID"`
	COMM_MODE  Comm_Mode      `json:"COMM_MODE"`
	ETHERNET   ETH_ST         `json:"ETHERNET"`
	TIMEOUT    TIMEOUT_ST     `json:"TIMEOUT"`
	CHANNELS   []MC3E_CHANNEL `json:"CHANNELS"`
	ENABLED    bool           `json:"ENABLED"`
	USE_DEVICE bool           `json:"USE_DEVICE"`
}

type ETH_ST struct {
	IP          string `json:"IP"`
	PORT        int    `json:"PORT"`
	ETH_DEVICE  string `json:"ETH_DEVICE"`
	Local_IP    string `json:"LOCAL_IP"`
	Local_Port  int    `json:"LOCAL_PORT"`
	Station_Num string `json:"STATION_NUM"`
}

type SERIAL_ST struct {
	PORT      string `json:"PORT"`
	BAUD      int    `json:"BAUD"`
	DATA_BITS int    `json:"DATA_BITS"`
	STOP_BITS int    `json:"STOP_BITS"`
	PARITY    string `json:"PARITY"`
}

type TIMEOUT_ST struct {
	CON_TIMEOUT   string `json:"CON_TIMEOUT"` // 1000ms
	WRITE_TIMEOUT string `json:"WRITE_TIMEOUT"`
	READ_TIMEOUT  string `json:"READ_TIMEOUT"`
	IDLE_TIMEOUT  string `json:"IDLE_TIMEOUT"`
}

type MC3E_COMMAND int

const (
	COMMAND_READ MC3E_COMMAND = iota
	COMMAND_BITREAD
	COMMAND_WRITE
)

func (s MC3E_COMMAND) String() string {
	switch s {
	case COMMAND_READ:
		return "READ"
	case COMMAND_BITREAD:
		return "BITREAD"
	case COMMAND_WRITE:
		return "WRITE"
	default:
		return "Unknown"
	}
}

type MB_COMMAND int

const (
	// Discrete (Coils / Inputs)
	COMMAND_READ_COILS           MB_COMMAND = 0x01 // Read Coils
	COMMAND_READ_DISCRETE_INPUTS MB_COMMAND = 0x02 // Read Discrete Inputs

	// Registers
	COMMAND_READ_HOLDING_REGISTERS MB_COMMAND = 0x03 // Read Holding Registers
	COMMAND_READ_INPUT_REGISTERS   MB_COMMAND = 0x04 // Read Input Registers
	COMMAND_WRITE_SINGLE_COIL      MB_COMMAND = 0x05 // Write Single Coil
	COMMAND_WRITE_SINGLE_REGISTER  MB_COMMAND = 0x06 // Write Single Register

	// Serial-only / Diagnostics / Events (장치/전송에 따라 지원 여부 다름)
	COMMAND_READ_EXCEPTION_STATUS  MB_COMMAND = 0x07 // Read Exception Status (Serial only)
	COMMAND_DIAGNOSTICS            MB_COMMAND = 0x08 // Diagnostics (Sub-function)
	COMMAND_GET_COMM_EVENT_COUNTER MB_COMMAND = 0x0B // Get Comm Event Counter
	COMMAND_GET_COMM_EVENT_LOG     MB_COMMAND = 0x0C // Get Comm Event Log

	// Multiple write / mask / combo
	COMMAND_WRITE_MULTIPLE_COILS     MB_COMMAND = 0x0F // Write Multiple Coils
	COMMAND_WRITE_MULTIPLE_REGISTERS MB_COMMAND = 0x10 // Write Multiple Registers
	COMMAND_REPORT_SERVER_ID         MB_COMMAND = 0x11 // Report Server ID (Report Slave ID)
	COMMAND_READ_FILE_RECORD         MB_COMMAND = 0x14 // Read File Record
	COMMAND_WRITE_FILE_RECORD        MB_COMMAND = 0x15 // Write File Record
	COMMAND_MASK_WRITE_REGISTER      MB_COMMAND = 0x16 // Mask Write Register
	COMMAND_READ_WRITE_MULTIPLE_REGS MB_COMMAND = 0x17 // Read/Write Multiple Registers
	COMMAND_READ_FIFO_QUEUE          MB_COMMAND = 0x18 // Read FIFO Queue

	// Encapsulated (Device Identification 등)
	COMMAND_ENCAPSULATED_INTERFACE_TRANSPORT MB_COMMAND = 0x2B // Encapsulated Interface Transport (MEI: e.g., Read Device Identification)
)

func (s MB_COMMAND) String() string {
	switch s {
	// Discrete
	case COMMAND_READ_COILS:
		return "READ COILS"
	case COMMAND_READ_DISCRETE_INPUTS:
		return "READ DISCRETE INPUTS"

	// Registers
	case COMMAND_READ_HOLDING_REGISTERS:
		return "READ HOLDING REGISTERS"
	case COMMAND_READ_INPUT_REGISTERS:
		return "READ INPUT REGISTERS"
	case COMMAND_WRITE_SINGLE_COIL:
		return "WRITE SINGLE COIL"
	case COMMAND_WRITE_SINGLE_REGISTER:
		return "WRITE SINGLE REGISTER"

	// Diagnostics / Events
	case COMMAND_READ_EXCEPTION_STATUS:
		return "READ EXCEPTION STATUS"
	case COMMAND_DIAGNOSTICS:
		return "DIAGNOSTICS"
	case COMMAND_GET_COMM_EVENT_COUNTER:
		return "GET COMM EVENT COUNTER"
	case COMMAND_GET_COMM_EVENT_LOG:
		return "GET COMM EVENT LOG"

	// Multiple / Mask / Combo
	case COMMAND_WRITE_MULTIPLE_COILS:
		return "WRITE MULTIPLE COILS"
	case COMMAND_WRITE_MULTIPLE_REGISTERS:
		return "WRITE MULTIPLE REGISTERS"
	case COMMAND_REPORT_SERVER_ID:
		return "REPORT SERVER ID"
	case COMMAND_READ_FILE_RECORD:
		return "READ FILE RECORD"
	case COMMAND_WRITE_FILE_RECORD:
		return "WRITE FILE RECORD"
	case COMMAND_MASK_WRITE_REGISTER:
		return "MASK WRITE REGISTER"
	case COMMAND_READ_WRITE_MULTIPLE_REGS:
		return "READ/WRITE MULTIPLE REGISTERS"
	case COMMAND_READ_FIFO_QUEUE:
		return "READ FIFO QUEUE"

	// Encapsulated (MEI)
	case COMMAND_ENCAPSULATED_INTERFACE_TRANSPORT:
		return "ENCAPSULATED INTERFACE TRANSPORT (MEI)"

	default:
		return "Unknown"
	}
}

type MC_DEVICE_CODE string

const (
	MC_DEVICE_B = "B"
	MC_DEVICE_D = "D"
	MC_DEVICE_W = "W"
	MC_DEVICE_R = "R"
)

type MB_SERVER_REGISTRY int

const (
	MB_SERVER_HOLDINGREGISTRY MB_SERVER_REGISTRY = iota
	MB_SERVER_INPUTREGISTRY
	MB_SERVER_COIL
)

func (s MB_SERVER_REGISTRY) String() string {
	switch s {
	case MB_SERVER_HOLDINGREGISTRY:
		return "HOLDING REGISTRY"
	case MB_SERVER_INPUTREGISTRY:
		return "INPUT REGISTRY"
	case MB_SERVER_COIL:
		return "COIL"
	}
	return "Unknown"
}

type MC3E_CHANNEL struct {
	COMMAND           MC3E_COMMAND       `json:"COMMAND"`
	DEVICE_CODE       MC_DEVICE_CODE     `json:"DEVICE_CODE"`
	OFFSET            int64              `json:"OFFSET"`
	NUM_POINT         int64              `json:"NUM_POINT"`
	MBSERVER_START    int                `json:"MBSERVER_START"`
	MBSERVER_REGISTRY MB_SERVER_REGISTRY `json:"MBSERVER_REGISTRY"`
	ENABLED           bool               `json:"ENABLED"`
}

type MB_CHANNEL struct {
	SLAVE_ID          byte               `json:"SLAVE_ID"`
	COMMAND           MB_COMMAND         `json:"COMMAND"`
	OFFSET            int                `json:"OFFSET"`
	NUM_POINT         int                `json:"NUM_POINT"`
	MBSERVER_START    int                `json:"MBSERVER_START"`
	MBSERVER_REGISTRY MB_SERVER_REGISTRY `json:"MBSERVER_REGISTRY"`
	ISCONNECT         bool               `json:"ISCONNECT"`
	ENABLED           bool               `json:"ENABLED"`
}

func LoadJSONFile[T any](filePath string) ([]T, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var result []T
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func SaveJsonFile[T any](filename string, data []T) error {

	bData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, bData, 0644)
}

func HaveChannels[T any](v T) bool {
	switch xs := any(v).(type) {
	case MC3E_DEVICE:
		return len(xs.CHANNELS) > 0
	case MODBUS_DEVICE:
		return len(xs.CHANNELS) > 0
	default:
		return false
	}
}
