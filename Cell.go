package main

import "crypto/sha256"

// -----------------------------------------
// Cell: the basic unit of the blockchain
// -----------------------------------------
type Cell struct {
    PrevHash [32]byte
    Data string
    DataHash [32]byte
}

// -----------------------------------------
// GetData(): return the content of Data field
// -----------------------------------------
func (c *Cell) GetData() string {
    return c.Data
}

// -----------------------------------------
// SetData(): set the Data field
// -----------------------------------------
func (c *Cell) SetData(d string) {
    c.Data = d
    c.ComputeDataHash()
}

// -----------------------------------------
// Compute the hash of the Data field
// -----------------------------------------
func (c *Cell) ComputeDataHash () {
    c.DataHash = sha256.Sum256( []byte(c.Data) )
}
// -----------------------------------------
// return the hash of Data field
// -----------------------------------------
func (c *Cell) GetDataHash() [32]byte {
    return c.DataHash
}

// -----------------------------------------
// Return the hash of previous cell
// -----------------------------------------
func (c *Cell) GetPrevHash() [32]byte {
    return c.PrevHash
}

// -----------------------------------------
// Set the hash of previous cell
// -----------------------------------------
func (c *Cell) SetPrevHash(h [32]byte)  {
    c.PrevHash = h
}

