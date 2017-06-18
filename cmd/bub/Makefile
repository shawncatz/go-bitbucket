BINARY = bin/bub

build: $(BINARY)

$(BINARY):
	go build -o $(BINARY)

test:
	go test -cover $$(go list ./...)

clean:
	rm -rf $(BINARY)

run: clean build
	$(BINARY)

deps:
	cd .. && glide install

install: clean build
	cp $(BINARY) $(GOPATH)/bin
