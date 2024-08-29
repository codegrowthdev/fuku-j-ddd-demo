package grp_ccommand_handler

import (
	"fuku-j-ddd-demo/ledger"
	"testing"
)

type LedgerLedgers interface {
	GetByLedgerId() (*ledger.Ledger, error)
}

type CommandHandler struct {
	ledgers ledger.Ledgers
}

func (self *CommandHandler) RegisterAsset(cmd RegisterAssetCommand) error {
	ledger, err := self.ledgers.GetByLedgerId(cmd.ledgerId)
	if err != nil {
		return err
	}

	return ledger.TryRegisterAsset(NewTryRegisterAsset())
}

func (ch *CommandHandler) AssetConnectBank() error {
	// Implementation goes here
	return nil
}

func (ch *CommandHandler) UpdateAsset() error {
	// Implementation goes here
	return nil
}

func (ch *CommandHandler) AskAsset() error {
	// Implementation goes here
	return nil
}

func (ch *CommandHandler) BidAsset() error {
	// Implementation goes here
	return nil
}

func (ch *CommandHandler) RegisterOwner() error {
	// Implementation goes here
	return nil
}

func TestCommandHandler(t *testing.T) {
	handler := CommandHandler{}

	t.Run("[Given] 買い手がいる時, [When] 買い手がマーケットに対して有効な買い注文をすると [Then] 買い手の元帳に記録される", func(t *testing.T) {
		// Given
		handler.AskAsset()
		buyer := Party.NewBuyer()
		ledger := Ledger.NewLedger(buyer)
		asset := Asset.NewAsset()
		transaction := Transaction.NewTransction()
		transaction = transaction.AddAsset(asset)
		transaction = transaction.Bit()

		// When
		transaction, result := transaction.TryAskedBy(buyer)
		assertSuccess(result)

		// Then
		_, result := ledger.TryAddAssetOfTransaction(transaction)
		assertSuccess(result)
	})

	t.Run("[Given] 買い手がいる時, [When] 買い手がマーケットに対して無効な買い注文をすると [Then] 注文は失敗する", func(t *testing.T) {
		// Given
		buyer := Party.NewBuyer()
		ledger := Ledger.NewLedger(buyer)
		asset := Asset.NewAsset()
		transaction := Transaction.NewTransction()
		transaction = transaction.AddAsset(asset)
		transaction = transaction.Bit()

		// When
		transaction, result := transaction.TryAskedBy(buyer)
		assertFail(result)

		// Then
		_, result := ledger.TryAddAssetOfTransaction(transaction)
		assertFail(result)
	})

	t.Run("RegisterAsset should throws error when brabrabra", func(t *testing.T) {
		// Test implementation goes here
	})

	t.Run("RegisterAsset should throws error when brabrabra", func(t *testing.T) {
		// Test implementation goes here
	})

	t.Run("AssetConnectBank should be work", func(t *testing.T) {
		// Test implementation goes here
	})

	t.Run("UpdateAsset should be work", func(t *testing.T) {
		// Test implementation goes here
	})

	t.Run("AskAsset should be work", func(t *testing.T) {
		// Test implementation goes here
	})

	t.Run("BidAsset should be work", func(t *testing.T) {
		// Test implementation goes here
	})

	t.Run("RegisterOwner should be work", func(t *testing.T) {
		// Test implementation goes here
	})
}
