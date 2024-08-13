package apigratis

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type Service struct {
    Server string
}

func NewService() *Service {
    return &Service{
        Server: "https://gateway.apibrasil.io/api/v2/",
    }
}

func (s *Service) Request(service string, dados string) (map[string]interface{}, error) {
    var data map[string]interface{}
    err := json.Unmarshal([]byte(dados), &data)
    if err != nil {
        return nil, err
    }

    credentials, ok := data["credentials"].(map[string]interface{})
    if !ok {
        return nil, fmt.Errorf("invalid request, missing credentials")
    }
    body, ok := data["body"].(map[string]interface{})
    if !ok {
        return nil, fmt.Errorf("invalid request, missing body")
    }

    payload, err := json.Marshal(body)
    if err != nil {
        return nil, err
    }

    url := s.Server + service

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
    if err != nil {
        return nil, err
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+credentials["BearerToken"].(string))
    req.Header.Set("DeviceToken", credentials["DeviceToken"].(string))
    req.Header.Set("User-Agent", "APIBRASIL/GOLANG-SDK")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var result map[string]interface{}
    err = json.Unmarshal(bodyBytes, &result)
    if err != nil {
        return nil, err
    }

    return result, nil
}

// whatsapp
func (s *Service) Whatsapp(dados string) (map[string]interface{}, error) {
    return s.Request("whatsapp", dados)
}

// sms
func (s *Service) Sms(dados string) (map[string]interface{}, error) {
	return s.Request("sms", dados)
}

// cpf/dados
func (s *Service) Cpf(dados string) (map[string]interface{}, error) {
	return s.Request("cpf/dados", dados)
}

//cnpj
func (s *Service) Cnpj(dados string) (map[string]interface{}, error) {
	return s.Request("dados", dados)
}