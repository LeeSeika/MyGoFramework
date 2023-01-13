package contract

const (
	EnvProduction  = "production"
	EnvTest        = "test"
	EnvDevelopment = "development"
	EnvKey         = "hade:env"
)

type Env interface {
	AppEnv() string
	IsExist(string) bool
	Get(string) string
	All() map[string]string
}
