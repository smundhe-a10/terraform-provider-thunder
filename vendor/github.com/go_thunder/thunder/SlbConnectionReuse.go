package go_thunder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"util"
)

type ConnectionReuse struct {
	UUID SlbConnectionReuseInstance `json:"connection-reuse,omitempty"`
}

type SamplingEnable7 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type SlbConnectionReuseInstance struct {
	Counters1 []SamplingEnable7 `json:"sampling-enable,omitempty"`
	UUID      string            `json:"uuid,omitempty"`
}

func GetSlbConnectionReuse(id string, host string) (*ConnectionReuse, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/connection-reuse", nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ConnectionReuse
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return nil, err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GetSlbConnectionReuse REQ RES..........................", m)
			check_api_status("GetSlbConnectionReuse", data)
			return &m, nil
		}
	}
}

func PostSlbConnectionReuse(id string, vc ConnectionReuse, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	payloadBytes, err := json.Marshal(vc)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/connection-reuse", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var v ConnectionReuse
		erro := json.Unmarshal(data, &v)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
			check_api_status("PostSlbConnectionReuse", data)
		}
	}

}
