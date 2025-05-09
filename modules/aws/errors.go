package aws

import "fmt"

type CacheNotFoundError struct {
	Hash string
}

func (e *CacheNotFoundError) Error() string {
	return fmt.Sprintf("Key %s Not Found Error", e.Hash)
}