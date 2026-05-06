package kafka

type ProducerConfig struct {
	Brokers      []string `mapstructure:"brokers"`
	Version      string   `mapstructure:"version"`
	SaslEnable   bool     `mapstructure:"sasl_enable"`
	SaslUser     string   `mapstructure:"sasl_user"`
	SaslPassword string   `mapstructure:"sasl_password"`
	TlsEnable    bool     `mapstructure:"tls_enable"`
}
