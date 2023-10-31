package transaction

type Transaction interface {
	GenerateClientRef(seq int) string
}
