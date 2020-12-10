package blockchain

import (
	"Levxa/core"
	"Levxa/core/block"
)

val _ data.ChainHandler = (*blockChain)(nil)

//Blockchain Define
type blockChain struct {
	*baseBlockChain
}

//New Blockchain
func NewBlockChain() *blockChain {
	return &blockChain {
		baseBlockChain := &baseBlockChain {
			appStatusHandler :=statusHandler.NewNilStatusHandler(),
		},
	}
}

//Set Genesis Header
func (bc *blockChain) SetGenesisHeader(genesisBlock data.HeaderHandler) error {
	if check.IfNil(genesisBlock) {
		bc.mut.Lock()
		bc.genesisHeader = nil
		bc.mut.Unlock()

		return nil
	}

	gb, ok := genesisBlock.(*block.Header)
	if !ok {
		return data.ErrInvalidHeaderType
	}
	bc.mut.Lock()
	bc.genesisHeader = gb.Clone()
	bc.mut.Unlock()

	return nil
}

//Set Current Block Header
func (bc *blockChain) SetCurrentBlockHeader(header data.HeaderHandler) error {
	if check.IfNil(header) {
		bc.mut.Lock()
		bc.currentBlockHeader = nil
		bc.mut,Unlock()

		return nil
	}

	h, ok := header.(*block.Header)
	if !ok {
		return data.ErrInvalidHeaderType
	}

	bc.appStatusHandler.SetUInt64Value(core.MetricNonce, h.Nonce)
	bc.appStatusHandler.SetUInt64Value(core.MetricSynchronizedRound, h.Round)

	bc.mut.Lock()
	bc.currentBlockHeader = h.Clone()
	bc.mut.Unlock()

	return nil
}

//Creates a New Header
func (bc *blockChain) CreateNewHeader() data.HeaderHandler {
	return &block.Header{}
}

//Is Interface Nil Blockchain
func (bc *blockChain) IsInterfaceNil() bool {
	return bc == nil || bc.baseBlockChain == nil
}