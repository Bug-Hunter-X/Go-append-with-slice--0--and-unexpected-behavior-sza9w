```go
func main() {

        x := make([]int, 0)

        x = append(x, 1)

        fmt.Println(x) // Output: [1]

        // Correct way to append to the existing slice
        x = append(x, 2)

        fmt.Println(x) // Output: [1 2]

        //Alternative to pre-allocate memory to the slice
        y := make([]int, 0, 10) // pre-allocate the space for 10 elements
        y = append(y, 3)
        fmt.Println(y)
        y = append(y, 4,5,6,7,8,9,10) //Append elements to the pre-allocated space
        fmt.Println(y)

}```