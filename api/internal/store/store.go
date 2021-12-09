package store

import (
    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
)

type DBConfig struct {
    Host string
    Port string
    User string
    Password string
    DBname string
}

type Store struct {
    config *DBConfig
    db *sql.DB
}

func New (config *DBConfig) *Store {
    return &Store {config: config}
}

func (s *Store) Open () error {
    psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
    s.config.Host, s.config.Port, s.config.User, s.config.Password, s.config.DBname)
    db, err := sql.Open("postgres", psqlconn)
    if err != nil {
        return err
    }
    if err := db.Ping(); err != nil {
        return err
    }
    s.db = db
    return nil
}

func (s *Store) Close (){
    s.db.Close();
}