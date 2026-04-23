package msconfig

var (
	XClientHost          string = "X-Client-Host"
	XClientCloudService  string = "X-Client-CloudService"
	XClientCloudInstance string = "X-Client-CloudInstance"

	// TODO: develop flow to proxy these headers
	XEndClientAppInterface string = "X-EndClient-AppInterface"
	XEndClientAppSlug      string = "X-EndClient-AppSlug"
)

type AppInterface string

var (
	AI_WebUI = "web_ui"
	AI_CLI   = "cli"
)
