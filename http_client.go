package utils

import (
    "fmt"
    "strings"
    "net/http"
    "io/ioutil"
)

var (
    bodyType = "application/json"
    //bodyType = "application/x-www-form-urlencoded"
)

func GET(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }

    if resp == nil {
        return nil, fmt.Errorf("resp is nil")
    }

    if resp.Body == nil {
        return nil, fmt.Errorf("body is nil")
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    return body, nil
}

func POST(url, data string) ([]byte, error) {
    resp, err := http.Post(url, bodyType, strings.NewReader(data))
    if err != nil {
        return nil, err
    }

    if resp == nil {
        return nil, fmt.Errorf("resp is nil")
    }

    if resp.Body == nil {
        return nil, fmt.Errorf("body is nil")
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    return body, nil
}

func DEL(url, data string) ([]byte, error) {
    req, err := http.NewRequest("DELETE", url, strings.NewReader(data))
    if err != nil {
        return nil, err
    }
    req.Header.Set("Content-Type", "application/json;charset=UTF-8")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }

    if resp == nil {
        return nil, fmt.Errorf("resp is nil")
    }

    if resp.Body == nil {
        return nil, fmt.Errorf("body is nil")
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    return body, nil
}
