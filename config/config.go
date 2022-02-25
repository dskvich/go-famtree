package config

type (
	Config struct {
		HTTP
		PG
	}

	HTTP struct {
		Port string `envconfig:"PORT" default:"8080"`
	}

	PG struct {
		URL     string `envconfig:"DATABASE_URL"`
		Host    string `envconfig:"DB_HOST" default:"localhost:65432"`
		ShowSQL bool   `envconfig:"DB_SHOW_SQL" default:"false"`
	}
)
