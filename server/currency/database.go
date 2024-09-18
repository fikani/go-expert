package currency

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type CotacaoRepository struct {
	db *sql.DB
}

func NewCotacaoRepository() (*CotacaoRepository, func() error) {
	// not the best place to open the database connection
	db, err := sql.Open("sqlite3", "cotacao.db")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS cotacao (id INTEGER PRIMARY KEY, code TEXT, codein TEXT, name TEXT, high TEXT, low TEXT, varBid TEXT, pctChange TEXT, bid TEXT, ask TEXT, timestamp TEXT, create_date TEXT)")
	if err != nil {
		panic(err)
	}
	return &CotacaoRepository{db: db}, db.Close
}

func (r *CotacaoRepository) Save(timeout time.Duration, cotacao CotacaoDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err := r.db.ExecContext(
		ctx,
		"INSERT INTO cotacao (code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		cotacao.USDBRL.Code,
		cotacao.USDBRL.Codein,
		cotacao.USDBRL.Name,
		cotacao.USDBRL.High,
		cotacao.USDBRL.Low,
		cotacao.USDBRL.VarBid,
		cotacao.USDBRL.PctChange,
		cotacao.USDBRL.Bid,
		cotacao.USDBRL.Ask,
		cotacao.USDBRL.Timestamp,
		cotacao.USDBRL.CreateDate,
	)
	if err != nil {
		fmt.Printf("Erro ao salvar cotação: %v\n", err)
	}
	return err
}
