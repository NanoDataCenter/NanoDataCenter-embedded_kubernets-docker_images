package main
 
import "fmt"
 
func main() {
    println("Simple Array:")
    var arrayint = [...]int{1, 2, 3, 4} //assign
    fmt.Println(arrayint, "\n")
 
    println("Simple Slice:")
    var sliceint []int
    sliceint = arrayint[:] //assign
    fmt.Println(sliceint, "\n")
 
    println("Array of arrays:")
    var arrayofarrays [3][len(arrayint)]int
    for i := range arrayofarrays { //assign
        arrayofarrays[i] = arrayint
    }
    fmt.Println(arrayofarrays, "\n")
 
    println("Array of slices:")
    var arrayofslice [len(arrayofarrays)][]int
    for i := range arrayofarrays { // assign
        arrayofslice[i] = arrayofarrays[i][:]
    }
    fmt.Println(arrayofslice, "\n")
 
    println("Slice of arrays:")
    var sliceofarray [][len(arrayint)]int
    sliceofarray = arrayofarrays[:]
    fmt.Println(sliceofarray, "\n")
 
    println("Slice of slices:")
    var sliceofslices [][]int
    sliceofslices = arrayofslice[:]
    fmt.Println(sliceofslices, "\n")
}