run:
	go install ./...
	emailapi

cert:
	mkcert localhost
	mv localhost-key.pem key.pem
	mv localhost.pem certificate.pem
