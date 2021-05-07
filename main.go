package main

import (
	"bufio"
	"embed"
	"log"
	"os"

	"github.com/candid82/joker/core"
	_ "github.com/candid82/joker/std/base64"
	_ "github.com/candid82/joker/std/bolt"
	_ "github.com/candid82/joker/std/crypto"
	_ "github.com/candid82/joker/std/csv"
	_ "github.com/candid82/joker/std/filepath"
	_ "github.com/candid82/joker/std/hex"
	_ "github.com/candid82/joker/std/html"
	_ "github.com/candid82/joker/std/http"
	_ "github.com/candid82/joker/std/io"
	_ "github.com/candid82/joker/std/json"
	_ "github.com/candid82/joker/std/markdown"
	_ "github.com/candid82/joker/std/math"
	_ "github.com/candid82/joker/std/os"
	_ "github.com/candid82/joker/std/strconv"
	_ "github.com/candid82/joker/std/string"
	_ "github.com/candid82/joker/std/time"
	_ "github.com/candid82/joker/std/url"
	_ "github.com/candid82/joker/std/uuid"
	_ "github.com/candid82/joker/std/yaml"
)

//go:embed src
var f embed.FS

func main() {
	var err error

	filename := "demo"

	core.GLOBAL_ENV.SetClassPath("src")
	core.GLOBAL_ENV.InitEnv(os.Stdin, os.Stdout, os.Stderr, os.Args[1:])
	core.GLOBAL_ENV.SetMainFilename(filename)
	core.GLOBAL_ENV.SetEmbedFs(f)
	core.ProcessCoreData()

	mainFile, err := f.Open("src/demo.joke")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	rdr := core.NewReader(
		bufio.NewReader(mainFile),
		filename,
	)
	err = core.ProcessReader(rdr, filename, core.EVAL)
	if err != nil {
		log.Fatal(err)
	}
}