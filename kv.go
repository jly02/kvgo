package store

import "fmt"

type Store struct {
	Data map[string]map[string]string
}

func New() *Store {
	s := &Store{
		Data: make(map[string]map[string]string),
	}

	return s
}

func (s *Store) Set(key string, field string, value string) {
	if s.Data[key] == nil {
		s.Data[key] = make(map[string]string)
	}

	s.Data[key][field] = value
}

func (s *Store) Get(key string, field string) (string, error) {
	record, ok := s.Data[key]
	if !ok {
		return "", fmt.Errorf("no such key %q", key)
	}

	value, ok := record[field]
	if !ok {
		return "", fmt.Errorf("no such field %q of key %q", field, key)
	}

	return value, nil
}

func (s *Store) Del(key string, field string) (bool, error) {
	if s.Data[key] == nil {
		return false, fmt.Errorf("no such key %q", key)
	}

	_, ok := s.Data[key][field]
	if !ok {
		return false, fmt.Errorf("no such field %q of key %q", field, key)
	}

	delete(s.Data, key)
	return true, nil
}
