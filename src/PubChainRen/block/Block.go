package block

import (
	"PubChainRen/utils"
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
	"PubChainRen/pow"
	"encoding/gob"
)



/**
 * 区块结构体
 */
type Block struct {
	//区块高度
	Height int64
	//上一个区块哈希
	PreHash []byte
	//交易数据
	Data []byte
	//时间戳
	TimeStamp int64
	//哈希
	Hash []byte
	//6.Nonce 随机值
	Nonce int64
}
func (block *Block) GetHeight() int64 {
	return block.Height
}

func (block *Block) GetPreHash() []byte {
	return block.PreHash
}

func (block *Block) GetData() []byte {
	return block.Data
}

func (block *Block) GetTimeStamp() int64 {
	return block.TimeStamp
}
func (block *Block) GetHash() []byte {
	return block.Hash
}
/**
 * 1.创建新的区块
 */
func NewBlock(data string, heigth int64, preHash []byte) *Block {
	//1.创建新区块
	block := &Block{Height: heigth, PreHash: preHash, Data: []byte(data), TimeStamp: time.Now().Unix(), Hash: nil}
	//2.调用工作量证明对象并返回符合要求的nonce值
	pow := pow.NewProofOfWork(block)
	hash, nonce := pow.Run()
	//2.设置block的hash
	block.Hash = hash;
	block.Nonce = nonce

	return block;
}

/**
 * 2.创建创世区块
 */
func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
/**
 * 把byte数据反序列化为区块Block结构
 */
func  (block *Block)Serialize() []byte  {
	var  result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err :=  encoder.Encode(block)
	if err != nil {
		panic(err.Error())
	}
	return  result.Bytes()

}
func  DeserializeBlock(blockByte []byte) *Block  {
	var  block Block
	decoder := gob.NewDecoder(bytes.NewReader(blockByte))
	err := decoder.Decode(&block)
	if err  != nil{
		panic(err.Error())
	}
	return &block
}
/**
 * 设置块结构hash
 */
func (block *Block) SetHash() {
	//1.height 转化为[]byte字节数据
	heightBytes := utils.IntToHex(block.Height)
	//2.时间戳转化为[]byte字节数组
	//第二个参数为进制2 ~ 36
	timeString := strconv.FormatInt(block.TimeStamp, 2)
	timeByte := []byte(timeString)
	//3.拼接所有属性
	blockBytes := bytes.Join([][]byte{heightBytes, block.PreHash, block.Data, timeByte, block.Hash}, []byte{})
	//4.生成hash
	hash := sha256.Sum256(blockBytes)
	//5.设置hash
	block.Hash = hash[:]

}
