all: build install

build:
	go build ./
	go build ./threads_locks/wiki

install:
	go install ./
	go install ./threads_locks/wiki



