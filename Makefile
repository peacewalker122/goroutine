hello:
	echo "Hello"
bench:
	go test -bench= -v 
test:
	go test -v ./... -cover