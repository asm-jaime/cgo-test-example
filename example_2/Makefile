all:
	gofmt -e -s -w .
	go vet .
	$(CC) -g -fPIC -c -o lib.o lib.c
	$(CC) -g -fPIC -shared -o liblib.so lib.o
	LD_lib_PATH=. go test
