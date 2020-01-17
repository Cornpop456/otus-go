package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var fromFlag string
var toFlag string
var offsetFlag int64
var limitFlag int64

// Copy copies data from one file to another with limit and offset passed through command line args
func Copy(from string, to string, limit int64, offset int64) error {
	src, err := os.Open(from)
	dst, err2 := os.Create(to)

	defer src.Close()
	defer dst.Close()

	if err != nil {
		return err
	} else if err2 != nil {
		return err2
	}

	fi, err := src.Stat()
	if err != nil {
		return err
	}

	srcSize := fi.Size()

	if srcSize-offset <= 0 {
		return fmt.Errorf("Incorrect offset: out of file size")
	} else if (srcSize-offset) < limit || limit == 0 {
		limit = srcSize - offset
	}

	src.Seek(offset, 0)

	tmpl := `{{green "Progress status:" }} {{percent . | magenta}} {{green "copied"}} `

	bar := pb.ProgressBarTemplate(tmpl).Start64(limit)

	barReader := bar.NewProxyReader(src)

	if _, err := io.CopyN(dst, barReader, limit); err != nil {
		return err
	}

	bar.Finish()
	return nil
}

func main() {
	flag.StringVar(&fromFlag, "from", "", "source file")
	flag.StringVar(&toFlag, "to", "", "dest file")
	flag.Int64Var(&offsetFlag, "offset", 0, "offset in src file")
	flag.Int64Var(&limitFlag, "limit", 0, "number of bytes to copy")

	flag.Parse()

	err := Copy(fromFlag, toFlag, limitFlag, offsetFlag)

	if err != nil {
		log.Fatal(err)
	}
}
