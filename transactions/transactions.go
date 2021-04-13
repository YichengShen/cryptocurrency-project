package transactions

// Transfer coins from one public key to another.
// TODO: Add coin ids.
type Transaction struct {
	signature string
	sender string
	recipient string
	coins int
}
