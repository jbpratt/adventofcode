open System
open System.IO

[<EntryPoint>]
let main argv =
    let calcSurfaceArea line = 
        let splitLine (line : string) = line.Split 'x'

        let calc (splitInts : string[]) = 
            2*3*Int32.Parse(splitInts.[0]) + 2*4*Int32.Parse(splitInts.[1]) + 2*2*Int32.Parse(splitInts.[2])
        
        calc
        // split line
    let data = File.ReadAllLines("input")

    0 // return an integer exit code