package database

import (
	//"database/sql"
	//"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/zenport.io/go-assignment/engine"
)

type Provider struct {
}

func (provider *Provider) GetKnightRepository() engine.FighterRepository {

	return &knightRepository{}
}

func (provider *Provider) Close() {

}

func NewProvider() *Provider {
	return &Provider{}
}
