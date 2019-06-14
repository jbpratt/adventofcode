open System
open System.IO

let splitLine (line : string) = 
    match line.Split 'x' |> Array.map System.Int64.Parse with
    |[|l; w; h|] -> (l,w,h)
    |_ -> raise <| System.Exception("Wrong input format")

[<EntryPoint>]
let main argv =
    let data = File.ReadAllLines("input") 

    let calcPaper (input : string[]) = 
        let calcPaperArea (l,w,h) = 
            let sides = [l*w; w*h; l*h]
            (List.sumBy ((*)2L) sides) + (List.min sides)

        input |> Seq.map splitLine |> Seq.sumBy calcPaperArea 
    
    let calcRibbon (input : string[]) =
        let calcRibbonLength (l,w,h) =
            let p = [2L*(l+w);2L*(w+h);2L*(l+h)]
            (List.min p) + (l*w*h)
        input |> Seq.map splitLine |> Seq.sumBy calcRibbonLength

    let p1 = data
            |> calcPaper 
            |> printfn "Part one %i"
    
    let p2 = data
            |> calcRibbon
            |> printfn "Part two %i"
    0