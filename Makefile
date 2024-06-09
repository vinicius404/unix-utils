build:
	go build -o dist/my-cat cmd/cat/main.go
	go build -o dist/my-ls cmd/ls/main.go
	go build -o dist/my-wc cmd/wc/main.go
	go build -o dist/my-grep cmd/grep/main.go

clean:
	go clean
	rm dist/my-cat
	rm dist/my-ls
	rm dist/my-wc
	rm dist/my-grep