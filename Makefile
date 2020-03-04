hash-n-go: worker master

gorilla:
	go get github.com/gorilla/websocket

worker: gorilla
	go install gitlab.com/hacheurs/hash-n-go/cmd/hash-n-go-worker

master: gorilla 
	go install gitlab.com/hacheurs/hash-n-go/cmd/hash-n-go

clean: 
	rm -rf bin