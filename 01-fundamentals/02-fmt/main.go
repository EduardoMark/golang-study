package main

import "fmt"

func main()  {
    fmt.Print("Same") 
    fmt.Print(" line.") 

    fmt.Println(" New")
    fmt.Println("line.")

    x := 17

    xs := fmt.Sprint(x) // retorna o conteúdo em string
    fmt.Println("O valor de x é:" + xs)
    fmt.Println("O valor de x é:", x)
    fmt.Printf("O valor de x é: %d\n", x)

    // %d (inteiro) %f (float) %t (bool) %s (string)
    fmt.Printf("%d %f %.2f %t %s %v", 17, 9.0, 81.99, true, "Hello", "World")
}