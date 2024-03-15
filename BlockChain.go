package main

import (
	"fmt"
    "bytes"
	"encoding/gob"
	"crypto/sha256"
)

// -------------------------------------
// Block: Estrutura base da blockchain
// -------------------------------------
type Block struct {
	prev *Block
	cell Cell
	next *Block
}

// --------------------------------------------
// BlockChain: Estrutura e metodos
// --------------------------------------------
type BlockChain struct {
	start *Block
	end   *Block
}

func (bc *BlockChain) serializeCell (c Cell, buff *bytes.Buffer) {
	// Cria o buffer o e encoder para a serialização
	//var buff bytes.Buffer
	enc := gob.NewEncoder(buff)
	enc.Encode(c)
	//serialized := buff.Bytes()
	//fmt.Println ("Estrutura serializada:", serialized)
}

// --------------------------------------------
// computeHash(): Calcula o hash de uma celula
// --------------------------------------------
func (bc *BlockChain) computeHash (cell Cell) [32]uint8 {
	// Serializa celula e calcula seu hash...
	var buff bytes.Buffer
	bc.serializeCell(cell, &buff)
	sc := buff.Bytes()
	h := sha256.Sum256( sc )

	return h
}


// --------------------------------------------
// BlockChain::Append() -
// --------------------------------------------
func (bc *BlockChain) Append (c Cell) {
	newBlock := &Block{cell: c}
	var pHash [32]uint8		// previous hash
	// Se BC vazia:
	if bc.start == nil {
		bc.start = newBlock
		// Como não existe um bloco anterior ao primeiro,
		// usaremos pHash vazio (todos os valores em 0).
		newBlock.cell.SetPrevHash(pHash)
	} else {
		newBlock.prev = bc.end
		bc.end.next = newBlock
		// Calcula o hash da celula anterior
		pHash = bc.computeHash(newBlock.prev.cell)
		newBlock.cell.SetPrevHash(pHash)
	}
	bc.end = newBlock
}

// --------------------------------------------
// BlockChain::Display() - Mostra o conteudo da Lista
// --------------------------------------------
func (bc *BlockChain) Display() {
	block := bc.start
	for block != nil {
		fmt.Println( "Data: ", block.cell.GetData() )
        fmt.Printf( "Data hash: %x\n", block.cell.GetDataHash() )
		fmt.Printf( "Prev hash: %x\n", block.cell.GetPrevHash() )

		block = block.next
	}
}

func (bc *BlockChain) validateChain() {
	block := bc.start
	count := 1
	var h [32]uint8
	var firstHash [32]uint8
	for block != nil {
		fmt.Printf ( "Bloco %d\n", count )
		fmt.Printf ( "Data hash: %x\n", block.cell.GetDataHash() )
		fmt.Printf ( "Prev hash: %x\n", block.cell.GetPrevHash() )
		if (block.prev == nil) {
			h = firstHash
		} else {
			h = bc.computeHash(block.prev.cell)
		}
		fmt.Printf ( "Calc hash: %x\n", h )
		if ( h == block.cell.GetPrevHash() ) {
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
