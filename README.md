# Magnacarto C API

Exposes a C API of [omniscale/magnacarto](https://github.com/omniscale/magnacarto) that can be called from other langugages such as Node.js or Python and builds Magnacarto into a shared library.

## API

`buildFromFile(file *C.char, options C.Opts) (*C.char)`

Expects a file path as well as a struct with options (see below) as input and returns a Mapnik or Mapserver style string.

`buildFromString(mmlStr *C.char, options C.Opts) (output, error *C.char)`

Expects a MML string in either JSON or YAML format and a struct with options (see below)
as input and returns a Mapnik or Mapserver style string.

The options struct has the following C type definition:

```
typedef struct {
    char* builderType;
    char* baseDir;
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
export $GOPATH=`pwd`
go get github.com/tools/godep
git clone https://github.com/gmgeo/libmagnacarto src/github.com/gmgeo/libmagnacarto
cd src/github.com/gmgeo/libmagnacarto
$GOPATH/bin/godep restore
make
```
