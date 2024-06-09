build:
	go build -o dist/my-cat cmd/cat/main.go
	go build -o dist/my-ls cmd/ls/main.go
	go build -o dist/my-wc cmd/wc/main.go

clean:
	go clean
	rm dist/my-cat
	rm dist/my-ls
	rm dist/my-wc