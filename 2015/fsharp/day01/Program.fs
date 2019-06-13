open System
open System.IO

[<EntryPoint>]
let main argv =
    let data = File.ReadAllText("input")
    let sum = ref 0
    for c in data do
        if c = '(' then
            incr sum
        else if c = ')' then
            decr sum

    printfn "%d" !sum
    0