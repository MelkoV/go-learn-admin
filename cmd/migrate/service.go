package migrate

import (
	"fmt"
	"github.com/MelkoV/go-learn-logger/logger"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"os"
)

func Up(l logger.CategoryLogger, path string, dsn string) {
	p := fmt.Sprintf("file://%s", path)
	l.Info("trying migrate from %s to %s", p, dsn)
	m := getConn(l, p, dsn)
	m.Up()
	l.Info("migrate up success")
}

func getConn(l logger.CategoryLogger, path, dsn string) *migrate.Migrate {
	m, err := migrate.New(path, dsn)
	if err != nil {
		l.Fatal("error db connection: %v", err)
		os.Exit(1)
	}
	return m
}
