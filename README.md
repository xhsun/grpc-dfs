# grpc-file-server
gRPC File Transfer Server


Generate CA private key and self-signed certificate
```
openssl req -x509 -newkey rsa:4096 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=CA/ST=STATE/L=CITY/O=DEV/OU=EXAMPLE/CN=FILETRANSFER/emailAddress=example@example.com"
```

Generate Web Server’s Private Key and CSR (Certificate Signing Request)
```
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=CA/ST=STATE/L=CITY/O=DEV/OU=EXAMPLE/CN=FILETRANSFER/emailAddress=example@example.com"
```