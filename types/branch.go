package types

type Branch struct {
	ID int32

	Name     string
	KeyStore []KeyStore

	COUNT int64
}
