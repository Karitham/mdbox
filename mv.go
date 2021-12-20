package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type mvParams struct {
	src      string
	dst      string
	pwd      string
	fileExts []string
}

var mdImageReg = regexp.MustCompile(`(?:\[(?P<text>.*?)\])\((?P<link>.*?)\)`)

func mv(p mvParams) error {
	filepath.WalkDir(p.pwd, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Println(err)
			return nil
		}
		// early exit
		if d.IsDir() || !elems(filepath.Ext(path), p.fileExts) {
			return nil
		}

		f, err := os.OpenFile(path, os.O_RDWR, d.Type())
		if err != nil {
			log.Println(err)
			return nil
		}
		defer f.Close()
		buf, err := io.ReadAll(f)
		if err != nil {
			return nil
		}

		buf = mdImageReg.ReplaceAllFunc(buf, func(b []byte) []byte {
			subs := mdImageReg.FindSubmatch(b)
			for i, s := range subs {
				if string(s) == p.src {
					subs[i] = []byte(p.dst)
				}
			}
			return []byte(fmt.Sprintf("[%s](%s)", string(subs[1]), string(subs[2])))
		})

		f.Seek(0, 0)
		f.Truncate(0)

		_, errW := f.Write(buf)
		return errW
	})

	return os.Rename(p.src, p.dst)
}
