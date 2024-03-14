package main

import "crypto/sha256"

type Cell struct {
    PrevHash [32]byte
    Data string
    DataHash [32]byte
}

func (c *Cell) getData() string {
    return c.Data
}

func (c *Cell) setData(d string) {
    c.Data = d
    c.computeDataHash()
}

func (c *Cell) computeDataHash () {
    c.DataHash = sha256.Sum256( []byte(c.Data) )
}

/*
func (c *Cell) computeHash (content Cell) {
    return sha256.Sum256( content )
}
*/

func (c *Cell) getDataHash() [32]byte {
    return c.DataHash
}

func (c *Cell) GetPrevHash() [32]byte {
    return c.PrevHash
}

func (c *Cell) SetPrevHash(h [32]byte)  {
    c.PrevHash = h
}