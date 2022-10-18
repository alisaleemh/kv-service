package service

import "errors"

func (s Service) GetKey(key string) ([]byte, error) {

	if key == "" {
		return nil, errors.New("key empty")
	}

	res := s.Storage.Db.Get(key)

	if len(res) == 0 {
		return nil, errors.New("cannot find key")
	}

	return res, nil
}

func (s Service) DeleteKey(key string) error {

	if key == "" {
		return errors.New("key empty")
	}

	return s.Storage.Db.Delete(key)
}

func (s Service) InsertKey(key string, value []byte) error {

	if key == "" {
		return errors.New("key empty")
	}

	if len(value) == 0 {
		return errors.New("value is empty")
	}

	return s.Storage.Db.Insert(key, value)
}
