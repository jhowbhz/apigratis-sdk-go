# SDK Go Lang - APIGratis by APIBrasil

Conjunto de API, para desenvolvedores.

_Transforme seus projetos em solucoes inteligentes com nossa API. Com recursos como API do WhatsApp, geolocalizacao, rastreamento de encomendas, verificacao de CPF/CNPJ e mais, voce pode criar solucoes eficientes e funcionais._

## Como instalar

```go get github.com/jhowbhz/apigratis-sdk-go```

## Canais de suporte (Comunidade)
[![WhatsApp Group](https://img.shields.io/badge/WhatsApp-Group-25D366?logo=whatsapp)](https://chat.whatsapp.com/EeAWALQb6Ga5oeTbG7DD2k) [![Telegram Group](https://img.shields.io/badge/Telegram-Group-32AFED?logo=telegram)](https://t.me/apigratisoficial)

## Obtenha suas credenciais
https://apibrasil.com.br

## Servicos de API disponiveis no SDK

| Up  | Services available            | Description       | Free    | Beta        | Stable   |
------|-------------------------------|-------------------|---------| ------------------------- | ------------------------- |
| Yes | WhatsAppService                | API do WhatsApp                         |   Yes   | Yes                   | Yes                   |
| Yes | SMS                            | API de SMS              .               |   Yes   | Yes                   | Yes                   |
| Yes | Receita Data CNPJ              | API Dados CNPJ Receita.                 |   No   | No                   | No                   |
| Yes | Receita Data CPF               | API Dados de CPF Serasa.                |   No   | No                   | No                   |
| Yes | CorreiosService                | API Busca encomendas Correios Brazil.   |   No   | No                   | No                   |
| Yes | CEPLocation                    | API CEP Geolocation + IBGE Brazil.      |   No   | No                   | No                   |
| Yes | VehiclesService                | API Placa Dados.                        |   No   | No                   | No                   |
| Yes | FipeService                    | API Placa FIPE.                         |   No   | No                   | No                   |


## Documentacoes
https://apibrasil.com.br/documentacoes

# SMS Service
```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/jhowbhz/apigratis-sdk-go/apigratis"
)

func main() {
    service := apigratis.NewService()

    // sendText
    send := map[string]interface{}{
        "action": "send",
        "credentials": map[string]interface{}{
            "DeviceToken": "YOUR_DEVICE_TOKEN",
            "BearerToken": "YOUR_BEARER_TOKEN",
		},
        "body": map[string]interface{}{
            "number": "5531994359434",
            "message": "Hello World for Go",
        },
    }

    payload, _ := json.Marshal(send)
    response, err := service.Sms(string(payload))

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Response:", response)
}
```

# WhatsApp Service
```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/jhowbhz/apigratis-sdk-go/apigratis"
)

func main() {

    service := apigratis.NewService()

    // sendText
    sendText := map[string]interface{}{
        "action": "sendText",
        "credentials": map[string]interface{}{
            "DeviceToken": "YOUR_DEVICE_TOKEN",
            "BearerToken": "YOUR_BEARER_TOKEN",
        },
        "body": map[string]interface{}{
            "text":        "Hello World for Go",
            "number":      "5531994359434",
            "time_typing": 1,
        },
    }

    payload, _ := json.Marshal(sendText)
    response, err := service.Whatsapp(string(payload))

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Response:", response)

    // sendFile
    sendFile := map[string]interface{}{
        "action": "sendFile",
        "credentials": map[string]interface{}{
            "DeviceToken": "YOUR_DEVICE_TOKEN",
            "BearerToken": "YOUR_BEARER_TOKEN",
        },
        "body": map[string]interface{}{
            "number" : "5531994359434",
            "path" : "https://assets.nagios.com/downloads/nagiosxi/docs/Installing_The_XI_Linux_Agent.pdf",
            "options" : {
                "caption": "texto do caption para arquivo",
                "createChat": True,
                "filename": "arquivo X"
            }
        },
    }

    payload, _ := json.Marshal(sendFile)
    response, err := service.Whatsapp(string(payload))

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Response:", response)
}
```