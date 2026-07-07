package main
import "fmt"
func main() {
    var mat [3][3]int

    for i := 0; i < 3; i++{
        for j := 0; j < 3; j ++{
            fmt.Scan(&mat[i][j])
        }
    }

    ref := mat[0][0] + mat[0][1] + mat[0][2]


    for i := 0; i < 3; i++{
        soma := 0
        for j := 0; j < 3; j++{
            soma += mat[i][j]
        }
        if soma != ref{
            fmt.Println("nao")
            return
        }
    }

    for j := 0; j < 3; j++{
        soma := 0
        for i := 0; i < 3; i++{
            soma += mat[i][j]
        }
        if soma != ref{
            fmt.Println("nao")
            return
        }
    }

    soma := 0

    for i := 0; i < 3; i++{
        soma += mat[i][i]
    }
    if soma != ref{
        fmt.Println("nao")
    }

    soma = 0

    for i := 0; i < 3; i++{
        soma += mat[i][2 - i]
    }

    if soma != ref{
        fmt.Println("nao")
    }
        fmt.Println("sim")

}

