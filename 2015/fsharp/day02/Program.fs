open System
open System.IO

[<EntryPoint>]
let main argv =
    let calcSurfaceArea (input : string[]) = 
        let splitLine (line : string) = 
            line.Split 'x' |> Array.map System.Int32.Parse

        let calc (present : int[]) = 
            2*present.[0]*present.[1] + 2*present.[1]*present.[2] + 2*present.[2]*present.[0]
        input |> Seq.map splitLine |> Seq.map calc |> Seq.sum 
    let data = File.ReadAllLines("input") |> calcSurfaceArea |> printfn "%i"
    0