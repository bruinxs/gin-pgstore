package ginpgstore

import (
	"database/sql"

	"github.com/antonlindstrom/pgstore"
	"github.com/gin-contrib/sessions"
	gsessions "github.com/gorilla/sessions"
)

//PGStore implement gin-contrib/sessions/store
type PGStore struct {
	*pgstore.PGStore
}

func NewPGStore(dbURL string, keyPairs ...[]byte) (*PGStore, error) {
	p, err := pgstore.NewPGStore(dbURL, keyPairs...)
	if err != nil {
		return nil, err
	}
	return &PGStore{PGStore: p}, nil
}

func NewPGStoreFromPool(db *sql.DB, keyPairs ...[]byte) (*PGStore, error) {
	p, err := pgstore.NewPGStoreFromPool(db, keyPairs...)
	if err != nil {
		return nil, err
	}
	return &PGStore{PGStore: p}, nil
}

func (p *PGStore) Options(options sessions.Options) {
	p.PGStore.Options = &gsessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}
