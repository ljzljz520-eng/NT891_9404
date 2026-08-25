package store

import (
	"encoding/json"
	"errors"
	"go.etcd.io/bbolt"
	"traveldeck/domain"
)

func (s *Store) SavePreference(p domain.UserPreference) error {
	b, e := json.Marshal(p)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error {
		bkt, e := tx.CreateBucketIfNotExists([]byte("preferences"))
		if e != nil {
			return e
		}
		return bkt.Put([]byte(p.ID), b)
	})
}
func (s *Store) LoadPreference(id string) (domain.UserPreference, error) {
	var p domain.UserPreference
	e := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("preferences"))
		if b == nil {
			return errors.New("preference not found")
		}
		v := b.Get([]byte(id))
		if v == nil {
			return errors.New("preference not found")
		}
		return json.Unmarshal(v, &p)
	})
	return p, e
}
func (s *Store) Count(bucket string) (int, error) {
	n := 0
	e := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return nil
		}
		return b.ForEach(func(_, v []byte) error {
			if v != nil {
				n++
			}
			return nil
		})
	})
	return n, e
}
