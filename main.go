package main

import "fmt"

func main() {
    bc := BlockChain{}

    var tr1 Transaction
    tr1.Init()
    tr1.AddInput("0001", 1.6)
    tr1.AddInput("0001", 1.2)
    h1 := tr1.computeHash()

    bc.Append(tr1)

    fmt.Printf("len=%d cap=%d %v\n", len(tr1.Inputs), cap(tr1.Inputs), tr1.Inputs)
    fmt.Printf("Total Input: %f\n", tr1.SumInputs())
    fmt.Printf("Transaction1 hash: %x\n", h1)

    var tr2 Transaction
    tr2.Init()
    tr2.AddInput("0001", 0.56)
    tr2.AddInput("0001", 0.25)
    h2 := tr2.computeHash()

    bc.Append (tr2)

    fmt.Printf("len=%d cap=%d %v\n", len(tr2.Inputs), cap(tr2.Inputs), tr2.Inputs)
    fmt.Printf("Total Input: %f\n", tr2.SumInputs())
    fmt.Printf("Transaction2 hash: %x\n", h2)

    var tr3 Transaction
    tr3.Init()
    tr3.AddInput("0001", 0.56)
    tr3.AddInput("0001", 0.25)
    bc.Append (tr3)

    bc.EvalChain()
}

