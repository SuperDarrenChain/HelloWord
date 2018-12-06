package pow

import (
	"PubChainRen/utils"
	"bytes"
	"crypto/sha256"
	"math/big"
)

/**
 * 256位二进制前面0的个数，也用来调节难度
 * 例如：
 *   256位二进制前面至少要有16个零
 */
const TargetBit = 16

//000085901496d018b33dbaa14dc1eccde3326479e8821005db6f40f748ea68ef

type ProofOfWork struct {
	Block BlockInterface //当前要验证的区块
	Target *big.Int //大数据存储
}
/**
 * block的接口
 */
type BlockInterface interface {
	GetHeight() int64
	GetPreHash() []byte
	GetData() []byte
	GetTimeStamp() int64
	GetHash() []byte
}

/**
 * 1.创建工作量证明实体
 */

/**
* 1.创建工作量证明实体
*/
func NewProofOfWork(block BlockInterface) *ProofOfWork {
	//1.big.Int对象 1

	//1.创建一个初始值为1的target
	target := big.NewInt(1)
	//2.左移256 - targetBit
	target = target.Lsh(target, 256-TargetBit)

	return &ProofOfWork{block, target}
}


/**
 * 2.执行run方法，对某一个区块进行工作量证明验证和计算，并返回符合要求的nonce值和hash
 */
func  (pow *ProofOfWork)Run() ([]byte , int64)  {
	//1.将block的属性拼接成字节数组
	//2.生成hash
	//3.判断hash是否有效，满足要求，停止循环
	var nonce int64
	nonce = 0
	var hash [32]byte
	var hashInt big.Int
	for {
		dataBytes := pow.PrepareData(nonce)
		hash = sha256.Sum256(dataBytes)
		hashInt.SetBytes(hash[:])
		if pow.Target.Cmp(&hashInt) == 1 {
			break
		}
		nonce ++

	}
	return hash[:],nonce
}
func (pow *ProofOfWork)IsValid() bool  {
	var  hashInt big.Int
	hashInt.SetBytes(pow.Block.GetHash())
	if pow.Target.Cmp(&hashInt) == 1 {
		return true
	}
	return false
}

/**
 * 准备hash计算数据
 */
func  (pow *ProofOfWork) PrepareData(nonce int64) []byte  {
blockBytes := bytes.Join([][]byte{
	utils.IntToHex(pow.Block.GetHeight()),
	pow.Block.GetPreHash(),
	pow.Block.GetData(),
	utils.IntToHex(pow.Block.GetTimeStamp()),
	utils.IntToHex(nonce)},[]byte{})
	return blockBytes
}