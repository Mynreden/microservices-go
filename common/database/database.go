package database

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"sync"
)

var dbPool *pgxpool.Pool
var once sync.Once

func GetDB(dataSourceName string) *pgxpool.Pool {
	once.Do(func() {
		var err error
		dbPool, err = pgxpool.Connect(context.Background(), dataSourceName)
		if err != nil {
			log.Fatalf("Unable to connect to database: %v\n", err)
		}
	})
	return dbPool
}

func CloseDB() {
	dbPool.Close()
}
