ssh-keygen: sshkeygen.go
	go build --ldflags '-extldflags "-static"' -o ssh-keygen sshkeyfingerprint.go

sshkeygen: sshkeygen.go
	go build --ldflags '-extldflags "-static"' -o sshkeygen sshkeygen.go

.PHONY: docker
docker: sshkeygen ssh-keygen
	docker build -t gogs .
