package main

import (
    "fmt"
)

type Buku struct {
    id       string
    judul    string
    penulis  string
    penerbit string
    eksemplar int
    tahun    int
    rating   int
}

type DaftarBuku []Buku

func insertionSortDescending(pustaka *DaftarBuku, n int) {
    for i := 1; i < n; i++ {
        temp := (*pustaka)[i]
        j := i
        for j > 0 && temp.rating > (*pustaka)[j-1].rating {
            (*pustaka)[j] = (*pustaka)[j-1]
            j--
        }
        (*pustaka)[j] = temp
    }
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
    if n > 0 {
        fmt.Println("Buku Terfavorit:", pustaka[0].judul, pustaka[0].penulis, pustaka[0].penerbit, pustaka[0].tahun)
    }
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
    count := 5
    if n < 5 {
        count = n
    }
    fmt.Println("5 Buku dengan Rating Tertinggi:")
    for i := 0; i < count; i++ {
        fmt.Println(pustaka[i].judul, pustaka[i].rating)
    }
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
    low, high := 0, n-1
    found := false
    for low <= high {
        mid := (low + high) / 2
        if pustaka[mid].rating == r {
            fmt.Println("Buku dengan Rating", r, "ditemukan:", pustaka[mid].judul)
            found = true
            break
        } else if pustaka[mid].rating < r {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }
    if !found {
        fmt.Println("Tidak ada buku dengan rating seperti itu")
    }
}

func main() {
    pustaka := DaftarBuku{
        {"1", "Go Programming", "John", "Go Press", 5, 2023, 90},
        {"2", "Learning Algorithms", "Jane", "Tech Books", 3, 2021, 85},
        {"3", "Data Structures", "Alice", "Data House", 10, 2020, 95},
        {"4", "Advanced Go", "Bob", "Go Press", 2, 2022, 80},
        {"5", "Go for Beginners", "Charlie", "Learning Inc.", 7, 2022, 88},
    }
    n := len(pustaka)
    insertionSortDescending(&pustaka, n)
    CetakTerfavorit(pustaka, n)
    Cetak5Terbaru(pustaka, n)
    var ratingToFind int
    fmt.Println("Masukkan rating buku yang dicari:")
    fmt.Scan(&ratingToFind)
    CariBuku(pustaka, n, ratingToFind)
}
