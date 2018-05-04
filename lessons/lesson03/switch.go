
func main() {
    n,_ := strconv.Atoi(os.Args[1])    

    switch n {
        case 1,2:
            fmt.Println("a")
        case 3:
            fmt.Println("b")
        default:
            fmt.Println("x")
    }

    switch {
        case n > 0:
            fmt.Println("positive")
        case n < 0:
            fmt.Println("negative")
        default:
            fmt.Println("0")
    }
}
