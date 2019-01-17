default: build

all: build

build: injuredfs

injuredfs:
	go generate
	GO111MODULE=on go build	

clean:
	rm -f injuredfs
