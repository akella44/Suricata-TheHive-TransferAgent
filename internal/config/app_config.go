package config

type AppConfig struct {
	RedisHostName  string `env:"REDIS_HOST, default=127.0.0.1:6379"`
	RedisPassword  string `env:"REDIS_PASSWORD, default=hjASdgnmQ2ier3"`
	ConsumeChannel string `env:"REDIS_CHANNEL, default=Suricata"`
	ApiAddr        string `env:"THEHIVE_API_HOST, default=127.0.0.1:9000"`
	ApiToken       string `env:"THEHIVE_API_TOKEN, default=T+bl4JnKd7+0jHaNIDcx//OcVq6AMZbY"`
}
