package jsonexporter

import (
	"encoding/hex"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/onflow/flow-go/ledger"
	"github.com/onflow/flow-go/ledger/complete"
	"github.com/onflow/flow-go/module/metrics"
)

// ExportLedger exports ledger key value pairs at the given blockID
func ExportLedger(ledgerPath string, targetstate string, outputPath string) error {

	led, err := complete.NewLedger(ledgerPath, complete.DefaultCacheSize, &metrics.NoopCollector{}, log.Logger, nil, 0)
	if err != nil {
		return fmt.Errorf("cannot create ledger from write-a-head logs and checkpoints: %w", err)
	}
	stateBytes, err := hex.DecodeString(targetstate)
	if err != nil {
		return fmt.Errorf("failed to decode hex code of state: %w", err)
	}
	var state ledger.State
	copy(state[:], stateBytes)
	err = led.DumpTrieAsJSON(state, outputPath)
	if err != nil {
		return fmt.Errorf("cannot dump trie as json: %w", err)
	}
	return nil
}
