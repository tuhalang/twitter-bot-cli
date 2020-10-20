package util

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"encoding/json"

	"github.com/joho/godotenv"
)

// Rule query in twitter
type Rule struct {
	ID    string `json:"id"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
}

// GetStreamData get stream data from twitter
func GetStreamData() {
	req, err := http.NewRequest("GET", "https://api.twitter.com/2/tweets/search/stream", nil)

	if err != nil {
		log.Fatal(err.Error())
	}

	token := os.Getenv("BEARER_TOKEN")
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("stream", "true")

	client := http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')

		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println(string(line))
	}
}

// DeleteAllRules delete all rules query
func DeleteAllRules() {

	rules := GetAllRules()

	payload := map[string](map[string][]string){
		"delete": {
			"ids": rules,
		},
	}

	jsonPayload, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "https://api.twitter.com/2/tweets/search/stream/rules", bytes.NewBuffer(jsonPayload))

	if err != nil {
		log.Fatal(err.Error())
	}

	token := os.Getenv("BEARER_TOKEN")
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log.Println(string(body))

}

// GetAllRules get all rules query
func GetAllRules() []string {
	req, err := http.NewRequest("GET", "https://api.twitter.com/2/tweets/search/stream/rules", nil)

	if err != nil {
		log.Fatal(err.Error())
	}

	token := os.Getenv("BEARER_TOKEN")
	req.Header.Add("Authorization", "Bearer "+token)

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var objmap map[string]json.RawMessage
	json.Unmarshal([]byte(string(body)), &objmap)

	data := objmap["data"]

	var rules []Rule

	err = json.Unmarshal(data, &rules)

	if err != nil {
		log.Fatal(err.Error())
	}

	var rulesStr []string
	for _, v := range rules {
		fmt.Println(v)
		rulesStr = append(rulesStr, v.ID)
	}

	return rulesStr
}

// AddRules add rules to query
func AddRules(rules []map[string]string) {
	payload := make(map[string][]map[string]string)

	payload["add"] = rules

	jsonPayload, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "https://api.twitter.com/2/tweets/search/stream/rules", bytes.NewBuffer(jsonPayload))

	if err != nil {
		log.Fatal(err.Error())
	}

	token := os.Getenv("BEARER_TOKEN")
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log.Println(string(body))
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot locate file .env !")
	}
	var rules = []map[string]string{
		{"value": "elector", "tag": "trump"},
	}
	DeleteAllRules()
	AddRules(rules)
	GetStreamData()
}
