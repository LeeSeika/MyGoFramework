package contract

const AppKey = "hade:app"

type App interface {
	Version() string
	BaseFolder() string
	ConfigFolder() string
	ProviderFolder() string
	MiddlewareFolder() string
	LogFolder() string
	CommandFolder() string
	RuntimeFolder() string
	TestFolder() string
	LoadAppConfig(map[string]string)
}
