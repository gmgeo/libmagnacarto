all: build/libmagnacarto.so

build/:
	mkdir build

build/libmagnacarto.so: *.go build
	$(GOPATH)/bin/godep go build -buildmode=c-shared -o libmagnacarto.so && mv libmagnacarto.so $(dir $@) && mv libmagnacarto.h $(dir $@)

clean:
	rm -rf build/
