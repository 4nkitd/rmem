package types

type KeyStore struct {
	ID int32

	Name string

	Header []string
	Body   []string

	headersCount int64
	COUNT        int64
}
