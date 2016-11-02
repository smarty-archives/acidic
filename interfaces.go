package kvacid

import (
	"io"
	"time"
)

type Storage interface {
	BeginTransaction(timeout time.Duration) Transaction
	ResumeTransaction(id string) (Transaction, error)

	Load(key string, etag string) (*LoadResult, error)
}

type Transaction interface {
	ID() string

	Store(*StoreRequest) error
	Load(key string, etag string) (*LoadResult, error)
	Delete(key string, etag string) error

	Commit() error
	Rollback() error
}

type StoreRequest struct {
	Key      string
	Value    io.Reader
	Metadata map[string]string
}

type LoadResult struct {
	Value    io.Reader
	Metadata map[string]string
}
