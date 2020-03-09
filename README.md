# Hash-n-go

Bruteforces a string up to 6 characters from the given MD5 hash.

## Execute

### Without docker

* Run the master : `go run cmd/hash-n-go/main.go -hash <hash> -workers <workers>`
* Run the worker(s) in a new terminal : `go run cmd/hash-n-go-worker/main.go <ws uri, by default ws://localhost:8080>`

### With docker

#### Worker

* Build the worker image : `docker build -t hash-n-go-worker --build-arg USERNAME=<gitlab username> --build-arg PASSWORD=<gitlab password>`
* Run the worker image : `docker run -e WS_URI=<websocket URI> hash-n-go-worker`

#### Master

* Build the master image : `docker build -t hash-n-go --build-arg USERNAME=<gitlab username> --build-arg PASSWORD=<gitlab password>`
* Run the master image : `docker run hash-n-go -hash <MANDATORY: hash> -workers <nb workers> -uri <websocket uri>`

