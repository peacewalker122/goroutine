hello:
	echo "Hello"
bench:
	go test -benchmem -run=^$ -coverprofile=/tmp/vscode-gopZPN0s/go-code-cover -bench . testfile
test:
	go test -v ./... -cover