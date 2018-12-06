package main

import (
	"PubChainRen/block"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	newBlock := block.NewBlock(" block data ", 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	fmt.Printf("区块hash：%x\n", newBlock.Hash)

	//序列化区块操作
	blockBytes := newBlock.Serialize()
	//反序列化操作
	deSeriazlizeBlock := block.DeserializeBlock(blockBytes)
	fmt.Printf("反序列化后的区块的hash：%x\n", deSeriazlizeBlock.Hash)
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// use dbname;
	db.Update(func(tx *bolt.Tx) error {
		//mysql 创建表  create table tableName
		bucket, err := tx.CreateBucket([]byte("block"))
		if err != nil {
			panic(err.Error())
		}
		//mysql:  插入数据 insert table values(...)
		err = bucket.Put([]byte("1"), []byte("davie"))
		if err != nil {
			panic(err.Error())
		}
		//mysql:查询数据: select * from table
		databytes := bucket.Get([]byte("1"))
		fmt.Println(databytes)

		return nil
	})

}
