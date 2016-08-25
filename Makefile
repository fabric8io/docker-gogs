.PHONY: docker
docker: build/ssh-hostkeygen build/ssh-keygen build/start-gogs Dockerfile
	docker build -t gogs .

build/ssh-keygen: cmds/ssh-keygen/sshkeyfingerprint.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build --ldflags '-extldflags "-static" -s -w' -o build/ssh-keygen cmds/ssh-keygen/sshkeyfingerprint.go

build/ssh-hostkeygen: cmds/ssh-hostkeygen/ssh-hostkeygen.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build --ldflags '-extldflags "-static" -s -w' -o build/ssh-hostkeygen cmds/ssh-hostkeygen/ssh-hostkeygen.go

build/start-gogs: cmds/start-gogs/start-gogs.go
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags="sqlite" --ldflags '-extldflags "-static" -s -w' -o build/start-gogs cmds/start-gogs/start-gogs.go

.PHONY: clean
clean:
	rm -rf build
