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

	t.Run("[Given] レッジャーが使える時, [When] the owner requests register a new asset with valid date, [Then] the asset should be registered", func(t *testing.T) {
		// Test implementation goes here
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
