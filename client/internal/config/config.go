package config

type Config struct {
	ProjectName        string `env-default:"file-transfer"`
	ServerAddress      string `env:"FILE_TRANSFER_CLI_SERVER_ADDR" env-default:"localhost:50051"`
	ServerHostOverride string `env:"FILE_TRANSFER_CLI_SERVER_HOST_OVERRIDE" env-default:"test.example.com"`
	CertFilePath       string `env:"FILE_TRANSFER_CLI_CERTIFICATE_PATH" env-default:"../../x509/ca-cert.pem"`
	FileChunkSize      int    `env-default:"65536"`
}
