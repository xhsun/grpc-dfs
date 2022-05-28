package config

type Config struct {
	ProjectName        string `env-default:"dfs-cli"`
	ServerAddress      string `env:"DFS_CLI_SERVER_ADDR" env-default:"localhost:50051"`
	ServerHostOverride string `env:"DFS_CLI_SERVER_HOST_OVERRIDE" env-default:"localhost"`
	CertFilePath       string `env:"DFS_CLI_CERTIFICATE_PATH" env-default:"../x509/ca-cert.pem"`
	FileChunkSize      int    `env-default:"65536"`
}
