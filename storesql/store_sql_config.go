package storesql

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type configSQL struct {
	usr string
	pwd string
	hst string
	prt string
	sch string
}

func (s configSQL) connStr() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s", s.usr, s.pwd, s.hst, s.prt, s.sch,
	)
}

func newConfigSQL(path string) *configSQL {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scnr := bufio.NewScanner(file)

	usr := scan(scnr, usrTag)
	pwd := scan(scnr, pwdTag)
	hst := scan(scnr, hstTag)
	prt := scan(scnr, prtTag)
	sch := scan(scnr, schTag)

	return &configSQL{usr: usr, pwd: pwd, hst: hst, prt: prt, sch: sch}
}

func scan(scnr *bufio.Scanner, tag string) string {
	scnr.Scan()
	return strings.TrimPrefix(scnr.Text(), tag)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const (
	usrTag = "user: "
	pwdTag = "password: "
	hstTag = "host: "
	prtTag = "port: "
	schTag = "schema: "
)
