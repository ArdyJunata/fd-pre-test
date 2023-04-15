package infra

import (
	"database/sql"
	infrahttp "fd-test/application/infra/http"
)

type Infra interface {
	Run()
}

type InfraBuilder struct{}

func NewInfraFactory() *InfraBuilder {
	return &InfraBuilder{}
}

func (i InfraBuilder) CreateInfraHttp(port string, db *sql.DB) (Infra, error) {
	return infrahttp.NewRouter(port, db), nil
}
