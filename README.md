# apigratis-sdk-go
 
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

	requestPayload := map[string]interface{}{
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

	payload, _ := json.Marshal(requestPayload)
	response, err := service.Whatsapp(string(payload))
	
    if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:", response)
}
```