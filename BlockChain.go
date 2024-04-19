package main

import (
	"fmt"

)

// -------------------------------------
// Block: Estrutura base da blockchain
// -------------------------------------
type Block struct {
	prev *Block
	Tx Transaction
	PrevTxHash [32]byte
	next *Block
}

// --------------------------------------------
// BlockChain: Estrutura e metodos
// --------------------------------------------
type BlockChain struct {
	start *Block
	end   *Block
}

/*
func (bc *BlockChain) serializeTx (t Transaction, buff *bytes.Buffer) {
	enc := gob.NewEncoder(buff)
	enc.Encode(t)
}
*/


// --------------------------------------------
// BlockChain::Append() -
// --------------------------------------------
func (bc *BlockChain) Append (t Transaction) {
	newBlock := &Block{Tx: t}
	// Se BC vazia:
	if bc.start == nil {
		bc.start = newBlock
	} else {
		newBlock.prev = bc.end
		bc.end.next = newBlock
		// Compute previous tx hash
		newBlock.PrevTxHash = newBlock.prev.Tx.computeHash()
	}
	bc.end = newBlock
}

// --------------------------------------------
// BlockChain::Display() - Mostra o conteudo da Lista
// --------------------------------------------
func (bc *BlockChain) Display() {
	block := bc.start
	for block != nil {
		fmt.Println( "Tx Id: ", block.Tx.GetId() )
		block = block.next
	}
}

func (bc *BlockChain) EvalChain() {
	block := bc.start
	count := 1
	var h [32]uint8
	var firstHash [32]uint8
	for block != nil {
		fmt.Printf ( "Bloco %d\n", count )
		fmt.Printf ( "UUID: %s\n", block.Tx.GetId() )
 		
		if (block.prev == nil) {
			h = firstHash
		} else {
			h = block.prev.Tx.computeHash()
		}

		fmt.Printf ( "Prev hash: %x\n", block.PrevTxHash )
		fmt.Printf ( "Calc hash: %x\n", h )

		if ( h == block.PrevTxHash ) {
			fmt.Println ("Previous hash OK")
		} else {
			fmt.Println ("Error: Previous hash mismatch.")
		}

		block = block.next
		count++
		fmt.Println("--")
	}
}

// REF: https://www.youtube.com/watch?v=hSeTG55WlLs
