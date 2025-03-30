package golalib

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/olachat/gola/ormtpl"
	"github.com/olachat/gola/structs"
)

type CodeGen struct {
	Driver string
}

func (g *CodeGen) GenTPL(t ormtpl.TplStruct, tplName string) []byte {
	buf := bytes.NewBufferString("")
	t.SetVersion(VERSION)
	err := ormtpl.GetTpl(g.Driver, tplName).Execute(buf, t)
	if err != nil {
		panic(t.GetName() + " " + tplName +
			" genTpl error:\n" + err.Error())
	}
	return buf.Bytes()
}

func (g *CodeGen) GenPackage(db *structs.DBInfo) map[string][]byte {
	files := make(map[string][]byte)

	genFiles := map[string]string{
		"02_package.gogo": db.Schema + "_goladb.go",
	}

	for genTpl, genPath := range genFiles {
		data, err := formatBuffer(g.GenTPL(db, genTpl))
		if err != nil {
			panic(db.Schema + " db code error:\n" + err.Error())
		}
		files[genPath] = data
	}

	return files
}

func (g *CodeGen) GenORM(t *structs.Table) map[string][]byte {
	files := make(map[string][]byte)

	tableFolder := t.Name + string(filepath.Separator)

	genFiles := map[string]string{
		"00_struct.gogo":     tableFolder + t.Name + ".go",
		"01_struct_idx.gogo": tableFolder + t.Name + "_idx.go",
	}

	for genTpl, genPath := range genFiles {
		data, err := formatBuffer(g.GenTPL(t, genTpl))
		if err != nil {
			panic(t.Name + " code error:\n" + err.Error())
		}
		files[genPath] = data
	}

	return files
}

var (
	rgxSyntaxError = regexp.MustCompile(`(\d+):\d+: `)
)

func formatBuffer(buf []byte) ([]byte, error) {
	output, err := format.Source(buf)
	if err == nil {
		return output, nil
	}

	matches := rgxSyntaxError.FindStringSubmatch(err.Error())
	if matches == nil {
		panic(errors.New("failed to format template: " + err.Error()))
	}

	lineNum, _ := strconv.Atoi(matches[1])
	scanner := bufio.NewScanner(bytes.NewReader(buf))
	errBuf := &bytes.Buffer{}
	line := 1
	for ; scanner.Scan(); line++ {
		if delta := line - lineNum; delta < -5 || delta > 5 {
			continue
		}

		if line == lineNum {
			errBuf.WriteString(">>>> ")
		} else {
			fmt.Fprintf(errBuf, "% 4d ", line)
		}
		errBuf.Write(scanner.Bytes())
		errBuf.WriteByte('\n')
	}

	return nil, fmt.Errorf("failed to format template\n\n%s", errBuf.Bytes())
}
