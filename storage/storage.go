package storage

import "database/sql"

//Instance of storage
type Storage struct {
	config *Config
	//Database File Descriptor
	db *sql.DB
}

//Storage constructor
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

//Open connection
func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.config.DatabaseURI)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	storage.db = db
	return nil
}

//Close connection
func (storage *Storage) Close() {
	storage.db.Close()
}
