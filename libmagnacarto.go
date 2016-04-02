package main

/*
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
*/
import (
    "C"
)

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"

    "github.com/omniscale/magnacarto/builder"
    "github.com/omniscale/magnacarto/builder/mapnik"
    "github.com/omniscale/magnacarto/builder/mapserver"
    "github.com/omniscale/magnacarto/config"
    "github.com/omniscale/magnacarto/mml"
)

//export buildFromString
func buildFromString(mmlStr *C.char, options C.Opts) (output, error *C.char) {
    mml := C.GoString(mmlStr)
    if mml == "" {
        return nil, C.CString("Please specify a MML string.")
    }
    baseDir := C.GoString(options.baseDir)
    if baseDir == "" {
        return nil, C.CString("Please specify a basedir path.")
    }

    return build(mml, baseDir, options)
}

//export buildFromFile
func buildFromFile(file *C.char, options C.Opts) (output, error *C.char) {
    filePath := C.GoString(file)
    if filePath == "" {
        return nil, C.CString("Please specify a file path.")
    }

    mml, err := ioutil.ReadFile(filePath)
    if err != nil {
        return nil, C.CString(fmt.Sprint("Error when reading from file ", filePath, ": ", err))
    }

    baseDir := C.GoString(options.baseDir)
    if baseDir == "" {
        baseDir = filepath.Dir(filePath)
    }

    return build(string(mml), baseDir, options)
}

func build(mmlStr string, baseDir string, options C.Opts) (output, error *C.char) {
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

    locator.SetBaseDir(baseDir)
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

    r := strings.NewReader(mmlStr)
    mmlData, err := mml.Parse(r)
    if err != nil {
        return nil, C.CString(fmt.Sprint(err))
    }

    var style bytes.Buffer
    for _, s := range mmlData.Stylesheets {
        if strings.HasSuffix(s, ".mss") {
            r, err := os.Open(filepath.Join(baseDir, s))
            if err != nil {
                return nil, C.CString(fmt.Sprint(err))
            }
            content, err := ioutil.ReadAll(r)
            if err != nil {
                return nil, C.CString(fmt.Sprint(err))
            }
            r.Close()
            style.Write(content)
        }
    }

    builder.BuildMapFromString(m, mmlData, style.String())

    var buf bytes.Buffer
    if err := m.Write(&buf); err != nil {
        return nil, C.CString(fmt.Sprint(err))
    }

    return C.CString(buf.String()), nil
}

func main() {
}
