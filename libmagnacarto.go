package main

/*
typedef struct {
    char* builderType;
    char* sqliteDir;
    char* fontDir;
    char* shapeDir;
    char* imageDir;
    _Bool relPaths;
    _Bool msNoMapBlock;
} Opts;
*/
import (
    "C"
)

import (
    "bytes"
    "fmt"
    "path/filepath"

    "github.com/omniscale/magnacarto/builder"
    "github.com/omniscale/magnacarto/builder/mapnik"
    "github.com/omniscale/magnacarto/builder/mapserver"
    "github.com/omniscale/magnacarto/config"
)

//export buildFromFile
func buildFromFile(file *C.char, options C.Opts) (output, error *C.char) {
    filePath := C.GoString(file)
    if filePath == "" {
        return nil, C.CString("Please specify a file path.")
    }

    conf := config.Magnacarto{}
    locator := conf.Locator()

    builderType := C.GoString(options.builderType)
    if builderType == "" {
        builderType = "mapnik2"
    }
    sqliteDir := C.GoString(options.sqliteDir)
	if sqliteDir != "" {
		conf.Datasources.SQLiteDirs = []string{sqliteDir}
	}
    fontDir := C.GoString(options.fontDir)
	if fontDir != "" {
		conf.Mapnik.FontDirs = []string{fontDir}
	}
    shapeDir := C.GoString(options.shapeDir)
	if shapeDir != "" {
		conf.Datasources.ShapefileDirs = []string{shapeDir}
	}
    imageDir := C.GoString(options.imageDir)
	if imageDir != "" {
		conf.Datasources.ImageDirs = []string{imageDir}
	}
    relPaths := bool(options.relPaths)
    if relPaths {
		locator.UseRelPaths(relPaths)
	}
    msNoMapBlock := bool(options.msNoMapBlock)

    locator.SetBaseDir(filepath.Dir(filePath))
    var m builder.MapWriter

    switch builderType {
    case "mapserver":
        mm := mapserver.New(locator)
        mm.SetNoMapBlock(msNoMapBlock)
        m = mm
    case "mapnik2":
        mm := mapnik.New(locator)
        mm.SetMapnik2(true)
        m = mm
    case "mapnik3":
        m = mapnik.New(locator)
    default:
        return nil, C.CString(fmt.Sprint("unknown builder ", builderType))
    }

    b := builder.New(m)
    b.SetMML(filePath)

    if err := b.Build(); err != nil {
        return nil, C.CString(fmt.Sprint(err))
    }

    var buf bytes.Buffer
    if err := m.Write(&buf); err != nil {
        return nil, C.CString(fmt.Sprint(err))
    }

    return C.CString(buf.String()), nil
}

func main() {
}
