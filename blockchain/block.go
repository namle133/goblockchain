package blockchain

type Block struct {
	PrevHash []byte
	Data     []byte
	Hash     []byte
	Nonce    int
}

type BlockChain struct {
	Blocks []*Block
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{prevHash, []byte(data), []byte{}, 0}
	pow := NewProof(block)

	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
func (chain *BlockChain) AddBlock(data string) {
	preBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, preBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
