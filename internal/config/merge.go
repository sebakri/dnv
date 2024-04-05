package config

func Merge(cfgs []*Config) (*Config, error) {
	merged := &Config{
		EnvironmentVariables: make(map[string]string),
	}

	for i := len(cfgs) - 1; i >= 0; i-- {
		for k, v := range cfgs[i].EnvironmentVariables {
			merged.EnvironmentVariables[k] = v
		}
	}

	return merged, nil
}
