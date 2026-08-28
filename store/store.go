package store

import (
	"encoding/json"
	"errors"
	"go.etcd.io/bbolt"
	"os"
	"traveldeck/domain"
)

var decksBucket = []byte("decks")
var drawsBucket = []byte("draws")

type Store struct {
	db   *bbolt.DB
	path string
}

func Open(path string) (*Store, error) {
	if path == "" {
		return nil, errors.New("path required")
	}
	db, e := bbolt.Open(path, 0600, nil)
	if e != nil {
		return nil, e
	}
	s := &Store{db: db, path: path}
	e = db.Update(func(tx *bbolt.Tx) error {
		if _, e := tx.CreateBucketIfNotExists(decksBucket); e != nil {
			return e
		}
		_, e = tx.CreateBucketIfNotExists(drawsBucket)
		return e
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	return s.db.Close()
}
func (s *Store) SaveDeck(d domain.Deck) error {
	b, e := json.Marshal(d)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(decksBucket).Put([]byte(d.ID), b) })
}
func (s *Store) LoadDeck(id string) (domain.Deck, error) {
	var d domain.Deck
	e := s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket(decksBucket).Get([]byte(id))
		if v == nil {
			return errors.New("deck not found")
		}
		return json.Unmarshal(v, &d)
	})
	return d, e
}
func (s *Store) SaveDraw(r domain.DrawRecord) error {
	b, e := json.Marshal(r)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(drawsBucket).Put([]byte(r.ID), b) })
}
func (s *Store) ListDraws(deck string) ([]domain.DrawRecord, error) {
	out := []domain.DrawRecord{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(drawsBucket).ForEach(func(_, v []byte) error {
			var r domain.DrawRecord
			if e := json.Unmarshal(v, &r); e != nil {
				return e
			}
			if deck == "" || r.DeckID == deck {
				out = append(out, r)
			}
			return nil
		})
	})
	return out, e
}
func Remove(path string) error { return os.Remove(path) }
