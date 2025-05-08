module go-app

go 1.21

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/go-resty/resty/v2 v2.10.0
	gotit_api_go_sdk v0.0.0
)

replace gotit_api_go_sdk => ./gotit-api-go-sdk 