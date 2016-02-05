all: build/libmagnacarto.so

build/:
	mkdir build

build/libmagnacarto.so: build
	cd build && \
	go build -buildmode=c-shared -o libmagnacarto.so ../libmagnacarto.go

crosscompile: build
	# requires xgo (https://github.com/karalabe/xgo)
	cd build && \
	$(GOPATH)/bin/xgo -go 1.6rc2 -buildmode=c-shared --targets=linux/amd64,linux/386,darwin/amd64,darwin/386 github.com/gmgeo/libmagnacarto

clean:
	rm -rf build/
