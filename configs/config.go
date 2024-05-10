package configs

type Conf struct {
	DBPath string `mapstructure:"DB_PATH"`
}

func LoadConfig(path string) (*Conf, error) {
	return &Conf{
		DBPath: "../../tmp/db",
	}, nil
}
