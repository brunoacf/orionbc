package main

import "fmt"

func main() {
    bc := BlockChain{}

    //c1 := Cell{"000", "data0", nil}
    var c1 = Cell {}
    var c2 = Cell {}
    var c3 = Cell {}
    var c4 = Cell {}
    var c5 = Cell {}
    
    c1.SetData ("Data1")
    c2.SetData ("Data2")
    c3.SetData ("Data3")
    c4.SetData ("Data4")
    c5.SetData ("Data5")
    
    bc.Append(c1)
    bc.Append(c2)
    bc.Append(c3)
    bc.Append(c4)
    bc.Append(c5)

    bc.validateChain()


    var tr1 Transaction
    tr1.Init()
    tr1.AddInput("0001", 1.6)
    tr1.AddInput("0001", 1.2)
    h1 := tr1.computeHash()

    fmt.Printf("len=%d cap=%d %v\n", len(tr1.Inputs), cap(tr1.Inputs), tr1.Inputs)
    fmt.Printf("Total Input: %f\n", tr1.SumInputs())
    fmt.Printf("Transaction1 hash: %x\n", h1)

    var tr2 Transaction
    tr2.Init()
    tr2.AddInput("0001", 0.56)
    tr2.AddInput("0001", 0.25)
    h2 := tr2.computeHash()

    fmt.Printf("len=%d cap=%d %v\n", len(tr2.Inputs), cap(tr2.Inputs), tr2.Inputs)
    fmt.Printf("Total Input: %f\n", tr2.SumInputs())
    fmt.Printf("Transaction2 hash: %x\n", h2)
}

