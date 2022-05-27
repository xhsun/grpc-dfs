package config

type Config struct {
	Port            uint16 `json:"port" env-default:"50051"`
	CertFilePath    string `json:"certFilePath" env:"FILE_TRANSFER_SERVER_CERTIFICATE" env-default:"../x509/server-cert.pem"`
	KeyFilePath     string `json:"keyFilePath" env:"FILE_TRANSFER_SERVER_KEY" env-default:"../x509/server-key.pem"`
	FileStoragePath string `json:"fileStoragePath" env-default:"serverFiles"`
	FileChunkSize   int    `env-default:"65536"`
}
