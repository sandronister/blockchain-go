package blockchain

import "fmt"

func (bc *BlockChain) ShowKeys() {
	keys, err := bc.Repository.GetKeys()

	if err != nil {
		panic(err)
	}

	for _, key := range keys {

		if key == "lh" {
			break
		}

		item, err := bc.Repository.GetBlock([]byte(key))

		if err != nil {
			fmt.Println(key)
			break
			//panic(err)
		}

		fmt.Println("---------------------------------------------------------------------------------------")
		fmt.Printf("Prev. hash: %x\n", item.PrevHash)
		fmt.Printf("Data: %s\n", item.Data)
		fmt.Printf("Hash: %x\n", item.Hash)
		fmt.Println("---------------------------------------------------------------------------------------")
	}
}
