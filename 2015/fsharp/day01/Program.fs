open System
open System.IO

[<EntryPoint>]
let main argv =
    let data = File.ReadAllText("input")
    let p1 = Seq.sumBy (fun c -> if c = '(' then 1 else -1) data
    let p2 = data 
                |> Seq.map(fun d -> if d = '(' then 1 else -1)
                |> Seq.scan (+) 0
                |> Seq.findIndex (fun f -> f = -1)
    printfn "Part one: %i" p1
    printfn "Part two: %i" p2
    0