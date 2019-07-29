package block

import (
	"blockchain/pkg/cryptoAPI"
	"time"
)

type Block struct {
	Index int
	Timestamp string
	Data int
	Hash string
	PrevHash string
}
var BlockChain []Block
func CreateBlock(oldBlock Block, data int) Block {

	B := Block{
		Index:     oldBlock.Index + 1,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  oldBlock.Hash,
	}

	B.Hash = cryptoAPI.GenerateSHA256Hash(B.PrevHash + B.Timestamp + string(B.Index) + string(data))

	return B

}

func IsValid(block Block, prevBlock Block) bool {

	if prevBlock.Index + 1 != block.Index {
		return false
	}

	if prevBlock.Hash != block.PrevHash {
		return false
	}

	if cryptoAPI.GenerateSHA256Hash(block.PrevHash + block.Timestamp + string(block.Index) + string(block.Data)) != block.Hash {
		return false
	}

	return true
}

func ReplaceChain(chain []Block) {
	if len(chain) > len(BlockChain) {
		BlockChain = chain
	}
}