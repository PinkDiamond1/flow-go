// (c) 2019 Dapper Labs - ALL RIGHTS RESERVED

package stdmap

import (
	"fmt"

	"github.com/dapperlabs/flow-go/model/flow"
)

// Blocks implements the blocks memory pool.
type Blocks struct {
	*Backend
}

// NewBlocks creates a new memory pool for blocks.
func NewBlocks() (*Blocks, error) {
	a := &Blocks{
		Backend: NewBackend(),
	}

	return a, nil
}

// Add adds an block to the mempool.
func (a *Blocks) Add(block *flow.Block) error {
	return a.Backend.Add(block)
}

// Get returns the block with the given ID from the mempool.
func (a *Blocks) Get(blockID flow.Identifier) (*flow.Block, error) {
	entity, err := a.Backend.Get(blockID)
	if err != nil {
		return nil, err
	}
	block, ok := entity.(*flow.Block)
	if !ok {
		panic(fmt.Sprintf("invalid entity in block pool (%T)", entity))
	}
	return block, nil
}

// All returns all blocks from the pool.
func (a *Blocks) All() []*flow.Block {
	entities := a.Backend.All()
	blocks := make([]*flow.Block, 0, len(entities))
	for _, entity := range entities {
		block, ok := entity.(*flow.Block)
		if !ok {
			panic(fmt.Sprintf("invalid entity in block pool (%T)", entity))
		}
		blocks = append(blocks, block)
	}
	return blocks
}
