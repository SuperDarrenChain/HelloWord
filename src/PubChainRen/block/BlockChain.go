package block

import "github.com/boltdb/bolt"

const DbName = "my.db"
const BlockTableName = "block"
const LastBlockByte = "lastHash"

/**
 *区块链结构体
 */
type BlockChain struct {
	Tip []byte   //最新的区块的hash值
	DB  *bolt.DB //存储区块数据的数据库
}

/**
 *1.创建带有创世区块的区块链
 */
func CreateBlockChainWithGenesisiBlock() *BlockChain {
	db, err := bolt.Open(DbName, 0600, nil)
	if err != nil {
		panic(err.Error())
	}
	var  blockHash []byte
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BlockTableName))
		if bucket != nil{

		}
	}
	
	genesBlock := CreateGenesisBlock("genesBlock")
	return &BlockChain{[]*Block{genesBlock}}
}

/**
 * 2.添加新区块到区块链中
 */
func (chain *BlockChain) AddBlockToBlockChain(data string, heigth int64, preHash []byte) {
	newBlock := NewBlock(data, heigth, preHash)
	chain.Block = append(chain.Block, newBlock)
}
