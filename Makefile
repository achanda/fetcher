.PHONY: build fmt test vet install

BIN_NAME=$(shell basename `pwd`)
CMD_PACKAGE=$(shell glide nv | grep cmd)

install:
	@go get -t -v $(glide nv)

build: 	vet test
	@go build -v -o ./bin/$(BIN_NAME) $(CMD_PACKAGE)
	@chmod +x ./bin/$(BIN_NAME)

fmt:
	@go fmt `glide nv`

vet:
	@go vet `glide nv`

test:
	@go test -cover -v -race `glide nv`

clean:
	@rm -rf bin
