package config

type Config struct {
	File                 string            `json:"file"`
	EnvironmentVariables map[string]string `json:"environment_variables"`
}
