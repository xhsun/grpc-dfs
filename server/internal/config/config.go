package config

type Config struct {
	Port            uint16 `json:"port" env:"DFS_SERVER_PORT" env-default:"50051"`
	CertFilePath    string `json:"certFilePath" env:"DFS_SERVER_CERTIFICATE" env-default:"../x509/server-cert.pem"`
	KeyFilePath     string `json:"keyFilePath" env:"DFS_SERVER_KEY" env-default:"../x509/server-key.pem"`
	FileStoragePath string `json:"fileStoragePath" env:"DFS_SERVER_STORAGE" env-default:"serverFiles"`
	FileChunkSize   int    `env-default:"65536"`
}
