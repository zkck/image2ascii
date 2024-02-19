package image2ascii

type Config struct {
	Color    bool
	AsciiMap string
}

func DefaultConfig() Config {
	return Config{
		Color:    true,
		AsciiMap: " .:-=+*#%@",
	}
}
