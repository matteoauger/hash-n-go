hash-n-go: worker master

worker: 
	go install gitlab.com/hacheurs/hash-n-go/cmd/hash-n-go-worker

master: 
	go install gitlab.com/hacheurs/hash-n-go/cmd/hash-n-go

clean: 
	rm -rf bin