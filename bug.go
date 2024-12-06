```go
func main() {


        x := make([]int, 0)

        x = append(x, 1)

        fmt.Println(x)

        x = append(x[:0], 2)

        fmt.Println(x)

}