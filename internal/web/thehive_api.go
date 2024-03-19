package web

import (
	"bytes"
	"encoding/json"
	"internal/model"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

func CreateTheHiveAlert(data model.SuricataMessageData, hostName string, apiToken string) {
	string_description, err := json.Marshal(data.HTTP)
	if err != nil {
		log.Default().Fatal("Error marshaling JSON:", err)
		return
	}

	requestBody, err := json.Marshal(map[string]interface{}{
		"type":         data.EventType,
		"source":       data.SrcIP,
		"sourceRef":    data.Host,
		"externalLink": data.DestIP,
		"title":        data.Alert.Category + " " + data.Timestamp,
		"description":  string(string_description),
		"severity":     data.Alert.Severity,
		"tags":         []string{data.Alert.Category, data.Alert.Signature, data.AppProto, data.Alert.Action},
		"flag":         true,
		"tlp":          0,
		"pap":          0,
		"summary":      "",
		"status":       "New",
		"assignee":     "Empty",
		"caseTemplate": "None",
	})
	if err != nil {
		log.Default().Fatal("Error creating JSON:", err)
		return
	}

	url := "http://" + hostName + "/api/alert"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Default().Fatal("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")

	dumpedRequest, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Default().Fatal("Error dumping request:", err)
		return
	}

	log.Default().Println("Request:", string(dumpedRequest))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Default().Fatal("Error sending request:", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Default().Fatal("Error reading response body:", err)
		return
	}
	log.Default().Println("Response Body:", string(body))
}
