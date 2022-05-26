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

Sign the Web Server Certificate Request (CSR)
```
openssl x509 -req -in server-req.pem -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem
```

See what is inside the certificate
```
openssl x509 -in server-cert.pem -noout -text
```

Generate new
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./filetransfer/filetransfer.proto
```
