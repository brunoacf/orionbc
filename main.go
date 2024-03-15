package main

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

}

