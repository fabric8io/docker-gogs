.PHONY: docker
docker: sshkeygen ssh-keygen
	docker build -t gogs .

ssh-keygen: sshkeyfingerprint.go
	CGO_ENABLED=0 go build --ldflags '-extldflags "-static" -s -w' -o ssh-keygen sshkeyfingerprint.go

sshkeygen: sshkeygen.go
	CGO_ENABLED=0 go build --ldflags '-extldflags "-static" -s -w' -o sshkeygen sshkeygen.go

.PHONY: clean
clean:
	rm -f ssh-keygen sshkeygen
