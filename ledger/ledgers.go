package ledger

type Ledgers interface {
	GetByLedgerId() (*Ledger, error)
}
