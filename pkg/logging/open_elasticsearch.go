package logging

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gofiber/fiber/v2"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

var esClient *elasticsearch.Client

func Config(a *fiber.App) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	fmt.Println("Hello")
	esClient = es

}

func CreateLog(p string, level string, data []byte) {
	message := p
	ctx := context.Background()
	logData := map[string]interface{}{
		"datetime":   time.Now(),
		"created_at": time.Now().Format("2006-01-02 15:04:05"),
		"message":    message,
		"severity":   level,
		"data":       string(data),
	}
	jsonString, err := convertMapToJSON(logData)
	if err != nil {
		log.Println("Error converting map to JSON:", err)
	}
	req := esapi.IndexRequest{
		Index:   "hrms",
		Body:    strings.NewReader(jsonString),
		Refresh: "true",
	}
	fmt.Println(jsonString)
	res, err := req.Do(ctx, esClient)
	fmt.Println(res)

	if err != nil {
		log.Fatalf("Error indexing document: %s", err)
	}
	defer res.Body.Close()
}

func convertMapToJSON(data map[string]interface{}) (string, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
