open System
open System.IO

[<EntryPoint>]
let main argv =
    let data = File.ReadAllText("input")

    let p1 = data
          |> Seq.map (fun c -> match c with | '>' -> (1,0) | '^' -> (0,1) | '<' -> (-1,0) | _ -> (0,-1))
          |> Seq.scan (fun (x1,y1) (x2,y2) -> (x1 + x2, y1 + y2)) (0,0)
          |> Seq.distinct
          |> Seq.length
          |> printfn "%i"
    0
