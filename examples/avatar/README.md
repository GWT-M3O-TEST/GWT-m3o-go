# Avatar

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/avatar/api](https://m3o.com/avatar/api).

Endpoints:

## Generate

Generate an unique avatar


[https://m3o.com/avatar/api#Generate](https://m3o.com/avatar/api#Generate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/avatar"
)

// Generate an unique avatar
func GenerateAvatarAndReturnBase64stringOfTheAvatar() {
	avatarService := avatar.NewAvatarService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := avatarService.Generate(&avatar.GenerateRequest{
		Format: "jpeg",
Upload: false,
Gender: "female",
Username: "",
	})
	fmt.Println(rsp, err)
	
}
```
## Generate

Generate an unique avatar


[https://m3o.com/avatar/api#Generate](https://m3o.com/avatar/api#Generate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/avatar"
)

// Generate an unique avatar
func GenerateAnAvatarAndUploadTheAvatarToMicrosCdn() {
	avatarService := avatar.NewAvatarService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := avatarService.Generate(&avatar.GenerateRequest{
		Gender: "female",
Username: "",
Format: "png",
Upload: true,
	})
	fmt.Println(rsp, err)
	
}
```
