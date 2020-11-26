package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type SnmpServerManagementIndex struct {
	MgmtIndex SnmpServerManagementIndexInstance `json:"management-index,omitempty"`
}

type SnmpServerManagementIndexInstance struct {
	MgmtIndex int    `json:"mgmt-index,omitempty"`
	UUID      string `json:"uuid,omitempty"`
}

func PostSnmpServerManagementIndex(id string, inst SnmpServerManagementIndex, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSnmpServerManagementIndex")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/snmp-server/management-index", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerManagementIndex
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			check_api_status("PostSnmpServerManagementIndex", data)

		}
	}

}

func GetSnmpServerManagementIndex(id string, host string) (*SnmpServerManagementIndex, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSnmpServerManagementIndex")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/snmp-server/management-index/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerManagementIndex
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			check_api_status("GetSnmpServerManagementIndex", data)
			return &m, nil
		}
	}

}
