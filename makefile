build:
	go build -o bin/qsw-cli

run: build
	./bin/qsw-cli

install: build
	sudo mv bin/qsw-cli /usr/local/bin/qsw-cli
