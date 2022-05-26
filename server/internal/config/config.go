package config

type Config struct {
	Port            uint16 `json:"port" env-default:"50051"`
	CertFilePath    string `json:"certFilePath" env-default:"../x509/server_cert.pem"`
	KeyFilePath     string `json:"keyFilePath" env-default:"../x509/server_key.pem"`
	FileStoragePath string `json:"fileStoragePath" env-default:"serverFiles/%"`
	FileChunkSize   int    `env-default:"65536"`
}
