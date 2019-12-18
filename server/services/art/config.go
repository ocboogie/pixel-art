package art

type Config struct {
	ArtSize   uint16
	ArtColors uint8
}

func DefaultConfig() Config {
	return Config{
		ArtSize:   25,
		ArtColors: 8,
	}
}
