# Sms

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/sms/api](https://m3o.com/sms/api).

Endpoints:

## Send

Send an SMS.


[https://m3o.com/sms/api#Send](https://m3o.com/sms/api#Send)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/sms"
)

// Send an SMS.
func SendSms() {
	smsService := sms.NewSmsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := smsService.Send(&sms.SendRequest{
		Message: "Hi there!",
To: "+447681129",
From: "Alice",
	})
	fmt.Println(rsp, err)
	
}
```
