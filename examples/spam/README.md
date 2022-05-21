# Spam

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/spam/api](https://m3o.com/spam/api).

Endpoints:

## Classify

Check whether an email is likely to be spam based on its attributes


[https://m3o.com/spam/api#Classify](https://m3o.com/spam/api#Classify)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/spam"
)

// Check whether an email is likely to be spam based on its attributes
func ClassifyAnEmail() {
	spamService := spam.NewSpamService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spamService.Classify(&spam.ClassifyRequest{
		Subject: "Welcome",
TextBody: "Hi there,\n\nWelcome to M3O.\n\nThanks\nM3O team",
From: "noreply@m3o.com",
To: "hello@example.com",
	})
	fmt.Println(rsp, err)
	
}
```
## Classify

Check whether an email is likely to be spam based on its attributes


[https://m3o.com/spam/api#Classify](https://m3o.com/spam/api#Classify)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/spam"
)

// Check whether an email is likely to be spam based on its attributes
func ClassifyAnEmailUsingTheRawData() {
	spamService := spam.NewSpamService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spamService.Classify(&spam.ClassifyRequest{
		EmailBody: "Subject: Welcome\r\nTo: hello@emaple.com\r\nFrom: noreply@m3o.com\r\n\r\nHi there,\n\nWelcome to M3O.\n\nThanks\nM3O team",
	})
	fmt.Println(rsp, err)
	
}
```
