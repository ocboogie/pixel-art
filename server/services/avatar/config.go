package avatar

type Config struct {
	Size   uint8
	Colors []string
}

func DefaultConfig() Config {
	return Config{
		Size: 5,
		Colors: []string{
			"#1abc9c", "#2ecc71", "#3498db", "#9b59b6", "#e74c3c",
		},
	}
}
