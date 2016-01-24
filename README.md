# Magnacarto C API

Exposes a C API of [omniscale/magnacarto](https://github.com/omniscale/magnacarto) that can be called from other langugages such as Node.js or Python and builds Magnacarto into a shared library.

## API

`buildFromFile(file *C.char, options C.Opts) (*C.char)`

Expects a file path as well as a struct with options as input and returns a Mapnik or Mapserver style string. The options struct has the following C type definition:

```
typedef struct {
    char* builderType;
    char* sqliteDir;
    char* fontDir;
    char* shapeDir;
    char* imageDir;
    _Bool relPaths;
    _Bool msNoMapBlock;
} Opts;
```
All these options are optional.

## Dependencies

* Go >= 1.5
* Godep
* C Toolchain

## Building

Let's assume you want to develop within the directory `libmagnacarto`, then do the following (Bash style):
```
mkdir libmagnacarto
cd libmagnacarto
export $GOROOT=`pwd`
go get github.com/tools/godep # add godep to your PATH for convenience
git clone https://github.com/gmgeo/libmagnacarto src/github.com/gmgeo/libmagnacarto
cd src/github.com/gmgeo/libmagnacarto
godep restore
make
```
