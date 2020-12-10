package blockchain

import (
	"sync"

	"Levxa/core"
)

type baseBlockChain struct {
	mut						sync.RWMutex
	appStatusHandler		core.appStatusHandler
	genesisHeader			data.HeaderHandler
	genesisHeaderHash		[]byte
	currentBlockHeader		data.HeaderHandler
	currentBlockHeaderHash	[]byte
}

//Set App Status Handler
func (bbc *baseBlockChain) SetAppStatusHandler(ash core.appStatusHandler) error {
	if check.IfNil(ash) {
		return ErrNilAppStatusHandler
	}

	bbc.mut.Lock()
	bbc.appStatusHandler = ash
	bbc.mut.Unlock()
	return nil
}

//Get Genesis Header
func (bbc *baseBlockChain) GetGenesisHeader() data.HeaderHandler {
	bbc.mut.RLock()
	defer bbc.mut.RUnlock()

	if check.IfNil(bbc.genesisHeader) {
		return nil
	}

	return bbc.genesisHeader.Clone()
}

//Get Genesis Header Hash
func (bbc *baseBlockChain) GetGenesisHeaderHash() []byte {
	bbc.mut.RLock()
	defer bbc.mut.RUnlock()

	return bbc.genesisHeaderHash
}

//Set Genesis Header Hash
func (bbc *baseBlockChain) SetGenesisHeaderHash(hash []byte) {
	bbc.mut.Lock()
	bbc.genesisHeaderHash = hash
	bbc.mut.Unlock()
}

//Get Current Block Header
func (bbc *baseBlockChain) GetCurrentBlockHeader() data.HeaderHandler {
	bbc.mut.RLock()
	defer bbc.mut.RUnlock()

	if check.IfNil(bbc.currentBlockHeader) {
		return nil
	}

	return bbc.currentBlockHeader.Clone()
}

//Get Current Block Header Hash
func (bbc *baseBlockChain) GetCurrentBlockHeaderHash() []byte {
	bbc.mut.RLock()
	defer bbc.mut.RUnlock()

	return bbc.currentBlockHeaderHash
}

//Set Current Block Header Hash
func (bbc *baseBlockChain) SetCurrentBlockHeaderHash(hash []byte) {
	bbc.mut.Lock()
	bbc.currentBlockHeaderHash = hash
	bbc.mut.Unlock()
}