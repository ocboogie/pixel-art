package art

type Config struct {
	Size   uint16
	Colors uint8
}

func DefaultConfig() Config {
	return Config{
		Size:   25,
		Colors: 8,
	}
}
