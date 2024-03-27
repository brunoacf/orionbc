package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"github.com/google/uuid"
)

// Note: Use go get github.com/google/uuid
// before compile or run.

// ========================================================
// Basic structure to hold records of transactions
// ========================================================
type Record struct {
	Addr string
	Ammount float64
}


// --------------------------------------------------------
// Record Setter: Address 
// --------------------------------------------------------
func (rec *Record) SetAddr (addr string) bool {
	rec.Addr = addr
	return true
}


// --------------------------------------------------------
// Record Setter: Ammount
// --------------------------------------------------------
func (rec *Record) SetAmmount (ammount float64) bool {
	rec.Ammount = ammount
	return true
}


// --------------------------------------------------------
// Record Getter: Address 
// --------------------------------------------------------
func (rec *Record) GetAddr () string {
	return rec.Addr
}


// --------------------------------------------------------
// Record Getter: Ammount
// --------------------------------------------------------
func (rec *Record) GetAmmount () float64 {
	return rec.Ammount
}



// ========================================================
// Structure of the transaction
// ========================================================
type Transaction struct {
	PrevHash [32]byte
	Id       string
	Inputs   []Record
	Outputs  []Record
	Sigs     []string
	ReqSigs  []string
	Hash     [32]byte
}


// --------------------------------------------------------
// Init():
// --------------------------------------------------------
func (tr *Transaction) Init () {
	tr.Id = uuid.New().String()
}

// --------------------------------------------------------
// SetId():
// --------------------------------------------------------
func (tr *Transaction) GetId () string {
	return tr.Id
}


// --------------------------------------------------------
// Serialize(): Serialize the transaction
// --------------------------------------------------------
func (tr *Transaction) Serialize (t *Transaction, buff *bytes.Buffer) {
	// Cria o buffer o e encoder para a serialização
	enc := gob.NewEncoder(buff)
	enc.Encode(t)
}


// --------------------------------------------------------
// computeHash(): Compute transaction hash
// --------------------------------------------------------
func (tr *Transaction) computeHash () [32]uint8 {
	// Serializa celula e calcula seu hash...
	var buff bytes.Buffer
	tr.Serialize(tr, &buff)
	srec := buff.Bytes()
	h := sha256.Sum256(srec)

	return h
}


// --------------------------------------------------------
// AddInput()
// --------------------------------------------------------
func (tr *Transaction) AddInput (addr string, ammount float64) bool {
	var rec Record
	rec.SetAddr (addr)
	rec.Ammount = ammount
	tr.Inputs = append(tr.Inputs, rec)
	//fmt.Printf("AddInput(): hash: %x\n ", tr.computeHash())
	// TODO: Adicionar assinatura de addr
	return true
}


// --------------------------------------------------------
// AddOutput()
// --------------------------------------------------------
func (tr *Transaction) AddOutput (addr string, ammount float64) bool {
	if ammount <= 0 {
		return false
	}
	var rec Record
	rec.SetAddr (addr)
	rec.SetAmmount(ammount) 
	tr.Outputs = append(tr.Outputs, rec)
	return true
}


// --------------------------------------------------------
// SumInputs()
// --------------------------------------------------------
func (tr *Transaction) SumInputs () float64 {
	sum := 0.0
	for _, v := range tr.Inputs {
		sum = sum + v.Ammount
	}
	return sum
}


// --------------------------------------------------------
// SumOutputs()
// --------------------------------------------------------
func (tr *Transaction) SumOutputs () float64 {
	sum := 0.0
	for _, v := range tr.Outputs {
		sum = sum + v.Ammount
	}
	return sum
}

// --------------------------------------------------------
// Check for overdraw
// --------------------------------------------------------
// Return: TRUE if overdrown detected
//         FALSE if overdrown not detected
// --------------------------------------------------------
func (tr *Transaction) Overdraw () bool {
	in := tr.SumInputs()
	out := tr.SumOutputs()
	if out >= in {
		// overdrown!
		return false
	} else {
		// not overdrown
		return true
	}
}

// --------------------------------------------------------
// checkInputs(): Search for 0 ou negative inputs
// --------------------------------------------------------
func (tr *Transaction) checkInputs() bool {
	for _, v := range tr.Inputs {
		if v.Ammount <= 0 {
			return false
		}
	}
	return true
}


// --------------------------------------------------------
// checkInputs(): Search for 0 ou negative outputs
// --------------------------------------------------------
func (tr *Transaction) checkOutputs() bool {
	for _, v := range tr.Outputs {
		if v.Ammount <= 0 {
			return false
		}
	}
	return true
}



// --------------------------------------------------------
// Validate(): Validation routine for transactions
// --------------------------------------------------------
func (tr *Transaction) Validate () bool {
	if tr.Overdraw() {
		// overdraw detected
		return false
	}
	if ! tr.checkInputs() {
		return false
	}
	if ! tr.checkOutputs() {
		return false
	}

	// TODO:
	// Check signatures
	// Check required signatures (escrow)

	return true
}