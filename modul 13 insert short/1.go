package main

import "fmt"

func insertionSortDescending(a []int) {
    n := len(a)
    for i := 1; i < n; i++ {
        temp := a[i]
        j := i
        for j > 0 && temp > a[j-1] {
            a[j] = a[j-1]
            j--
        }
        a[j] = temp
    }
}

func checkEqualSpacing(a []int) string {
    if len(a) < 2 {
        return "Data berjarak tetap"
    }
    diff := a[1] - a[0]
    for i := 2; i < len(a); i++ {
        if a[i]-a[i-1] != diff {
            return "Data berjarak tidak tetap"
        }
    }
    return fmt.Sprintf("Data berjarak %d", diff)
}

func main() {
    var numbers []int
    var input int
    fmt.Println("Masukkan data (akhiri dengan bilangan negatif):")
    for {
        fmt.Scan(&input)
        if input < 0 {
            break
        }
        numbers = append(numbers, input)
    }
    insertionSortDescending(numbers)
    fmt.Println("Data setelah diurutkan secara descending:")
    fmt.Println(numbers)
    result := checkEqualSpacing(numbers)
    fmt.Println(result)
}
