package main

//import "crypto/sha256"

// -----------------------------------------
// Transaction Cell: 
// Contains the hash of the previous transaction,
// the current transaction and the current hash.
// -----------------------------------------
type Cell struct {
    PrevHash [32]byte
    //Data string
    Trx Transaction
    TxHash [32]byte
}


// -----------------------------------------
// Compute the hash of the Data field
// -----------------------------------------
func (c *Cell) ComputeDataHash () {
    //s := c.Tx.Serialize()
    //c.TxHash = sha256.Sum256( []byte(c.Tx) )
    c.TxHash = c.Trx.computeHash()
}
// -----------------------------------------
// return the hash of Transaction field
// -----------------------------------------
func (c *Cell) GetTxHash() [32]byte {
    return c.TxHash
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

