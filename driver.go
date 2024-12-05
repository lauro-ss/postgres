package postgres

import (
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/olauro/goe"
)

type Driver struct {
	dns       string
	returning []byte
}

func Open(dns string) (driver *Driver) {
	return &Driver{
		dns:       dns,
		returning: []byte(" RETURNING "),
	}
}

func (dr *Driver) Init(db *goe.DB) {
	if db.ConnPool != nil {
		return
	}
	config, err := pgx.ParseConfig(dr.dns)
	if err != nil {
		//TODO: Add error handling
		fmt.Println(err)
		return
	}
	db.ConnPool = stdlib.OpenDB(*config)
}

func (dr *Driver) KeywordHandler(s string) string {
	return fmt.Sprintf(`"%s"`, s)
}

func (dr *Driver) Returning(b []byte) []byte {
	return append(dr.returning, append(b, ';')...)
}
