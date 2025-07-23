package main

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
)

type ApiResponse struct {
    // e.g. Field1 string `json:"field1"`
    //      Field2 int    `json:"field2"`
}

func fetchData(url string) (*ApiResponse, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var data ApiResponse
    if err := json.Unmarshal(body, &data); err != nil {
        return nil, err
    }
    return &data, nil
}

func main() {
    url := "https://api.example.com/data"
    data, err := fetchData(url)
    if err != nil {
        log.Fatalf("Fetch error: %v", err)
    }
    fmt.Printf("Fetched data: %+v\n", data)
}
