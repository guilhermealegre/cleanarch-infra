package postgres

import (
	"context"
	"fmt"
	"github.com/gocraft/dbr/v2"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"log"
)

type postgres struct {
	Conn   *dbr.Connection
	config Config
}

func New(config Config) *postgres {
	return &postgres{
		config: config,
	}
}

func (p *postgres) Init() error {

	dns := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		p.config.Host,
		p.config.Port,
		p.config.User,
		p.config.Password,
		p.config.Database,
	)

	//Open database connections
	conn, err := dbr.Open(p.config.Driver, dns, nil)
	if err != nil {
		log.Println("Error opening database: ", err)
		panic(err)
	}
	conn.SetMaxOpenConns(10)

	//Test connection
	err = conn.Ping()
	if err != nil {
		log.Println("Error pinging database: ", err)
		panic(err)
	}

	p.Conn = conn

	//Run database migrations
	migrationNum, err := p.runMigrations(p.config.Driver)
	if err != nil {
		return err
	}
	fmt.Println("Number of Database Migrations Run: ", migrationNum)

	return nil
}

func (p *postgres) Read(ctx context.Context) *dbr.Session {
	session := p.Conn.NewSession(nil)
	return session
}

func (p *postgres) Write(ctx context.Context) *dbr.Session {
	session := p.Conn.NewSession(nil)
	return session
}

func (p *postgres) runMigrations(drive string) (int, error) {

	migrations := &migrate.MigrationSet{
		SchemaName: p.config.Schema,
	}

	//Set migration directory
	migrationsDir := &migrate.FileMigrationSource{
		Dir: p.config.Migrations + "/postgres",
	}

	return migrations.Exec(p.Conn.DB, drive, migrationsDir, migrate.Up)
}
