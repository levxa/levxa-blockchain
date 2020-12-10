package block

import (
	"math/big"

	"Levxa/core"
)

//MiniBlockSlice
type MiniBlockSlice []*MiniBlock

//Set Nonce
func (h *Header) SetNonce(n uint64) {
	h.Nonce = n
}

//Set Epoch
func (h *Header) SetEpoch(e uint32) {
	h.Epoch = e
}

//Set Round
func (h *Header) SetRound(r uint64) {
	h.Round = r
}

//Set Root Hash
func (h *Header) SetRootHash(rHash []byte) {
	h.RootHash = rHash
}

//Set Validator Stats Root Hash
func (h *Header) SetValidatorStatsRootHash(_ []byte) {
}

//Get Validator Stats Root Hash
func (h *Header) GetValidatorStatsRootHash() []byte {
	return []byte{}
}

//Set Previous Hash
func (h *Header) SetPrevHash(pvHash []byte) {
	h.PrevHash = pvHash
}

//Set Previous Random Seed
func (h *Header) SetPrevRandSeed(pvRandSeed []byte) {
	h.PrevRandSeed = pvRandSeed
}

//Set Random Seed
func (h *Header) SetRandSeed(randSeed []byte) {
	h.RandSeed = randSeed
}

//Set Public Key Bitmap
func (h *Header) SetPubKeyBitmap(pkbm []byte) {
	h.PubKeyButmap = pkbm
}

//Set Signature
func (h *Header) SetSignature(sg []byte) {
	h.Signature = sg
}

//Set Leader Signature (RALATTTTTT)
func (h *Header) SetLeaderSignature(sg []byte) {
	h.LeaderSignature = sg
}

//Set Chain ID
func (h *Header) SetChainID(chainID []byte) {
	h.ChainID = chainID
}

//Set Software Version
func (h *Header) SetSoftwareVersion(version []byte) {
	h.SoftwareVersion = version
}

//Set TimeStamp
func (h *Header) SetTimeStamp(ts uint64) {
	h.TimeStamp = ts
}

//Set Accumulated Fees
func (h *Header) SetAccumulatedFees(value *big.Int) {
	h.AccumulatedFees.Set(value)
}

//Set Developer Fees
func (h *Header) SetDeveloperFees(value *big.Int) {
	h.DeveloperFees.Set(value)
}

//Set Transaction Count or Tx Count
func (h *Header) SetTxCount(txCount uint32) {
	h.TxCount = txCount
}

//Set Shard ID
func (h *Header) SetShardID(shId uint32) {
	h.ShardID = shId
}

//Get Mini Block Headers With Dst
func (h *Header) GetMiniBlockHeaderWithDst(destId uint32) map[string]uint32 {
	hashDst := make(map[string]uint32)
	for _, val := range h.MiniBlockHeaders {
		if val.ReceiverShardID == destId && val.SenderShardID != destId {
			hashDst[string(val.Hash)] = val.SenderShardID
		}
	}
	return hashDst
}

//Get Ordered Cross Mini Block With Dst
func (h *Header) GetOrderedCrossMiniBlocksWithDst(destId uint32) []*data.MiniBlockInfo {
	miniBlock := make([]*data.MiniBlockInfo, 0)

	for _, mb := range h.MiniBlockHeaders {
		if mb.ReceiverShardID == destId && mb.SenderShardID != destId {
			miniBlocks = append(miniBlocks, &data.MiniBlockInfo {
				Hash 			: mb.Hash,
				SenderShardID 	: mb.SenderShardID,
				Round 			: h.Round,
			})
		}
	}

	return miniBlocks
}

//Get Mini Block Headers Hashes
func (h *Header) GetMiniBlockHeadersHashes() [][]byte {
	result := make([][]byte, 0, len(h.MiniBlockHeaders))
	for _, MiniBlock := range h.MiniBlockHeaders {
		result = append(result, miniBlock.Hash)
	}
	return result
}

//Map Mini Block Hashes To Shard
func (h *Header) MapMiniBlockHashesToShard() map[string]uint32 {
	hashDst := make(map[string]uint32)
	for _, val := range h.MiniBlockHeaders {
		hashDst[string(val.Hash)] = val.SenderShardID
	}
	return hashDst
}

//Clone Return Object
func (h *Header) Clone() data.HeaderHandler {
	headerCopy := *h
	return &headerCopy
}

//Integrity And Validity
func (b *Body) IntegrityAndValidity() error {
	if b.IsInterfaceNil() {
		return data.ErrNilBlockBody
	}

	for i := 0; i < len(b.MiniBlocks); i++ {
		if len(b.MiniBlocks[i].TxHashes) == 0 {
			return data.ErrMiniBlockEmpty
		}
	}

	return nil
}

//Clone Return Object
func (b *Body) Clone() data.BodyHandler {
	bodyCopy := *b

	return &bodyCopy
}

//Is Interface nil Body
func (b *Body) IsInterfaceNil() bool {
	return b == nil
}

//Is Interface Nil Header
func (h *Header) IsInterfaceNil() bool {
	return h == nil
}

//Is Start Of Epoch Block
func (h *Header) IsStartOfEpochBlock() bool {
	return len(h.EpochStartMetaHash) > 0
}

//Clone The Underlying Data
func (mb *MiniBlock) Clone() *MiniBlock {
	newMb := &MiniBlock {
		ReceiverShardID	: mb.ReceiverShardID,
		SenderShardID	: mb.SenderShardID,
		Type 			: mb.Type,
	}
	newMb.TxHashes = make([][]byte, len(mb.TxHashes))
	copy(newMb.TxHashes, mb.TxHashes)

	return newMb
}