.PHONY: assets go_binary build docker
# Build go binary.
go_binary: 
	go build -o zypress .

# Target to build a release binary.
build: go_binary

# Build image and run zypress server (with default settings).
docker:
	docker build -t zypress .
	docker run --rm \
		--name zypress \
		-p 2001:2001 \
		zypress --redis-addr=host.docker.internal:6379