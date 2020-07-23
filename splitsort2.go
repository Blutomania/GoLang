package main

import (
    "fmt"
    "strconv"
    "sync"
    "sort"
)

// Calling main. It's a tight main which calls a buncha routines and GoRoutines. 
func main() { 
    var wg sync.WaitGroup
    slice := generateSlice()
    howl := len(slice)
    var segm int
    if (howl % 4 == 0) {
        segm = (howl/4)
    } else {
        segm = (howl/3)
    }
    ch := make(chan int, howl)
    for i := 0; i < howl; i = i + segm {
        count := 0
        if ((howl - i) < 4) {
            segm = (howl-i)
        }
        var sli []int = slice[i:(i+segm)]
        switch count {
        case 1:
            go sortIt1(sli, ch)
        case 2:
            go sortIt2(sli, ch)
        case 3:
            go sortIt3(sli, ch)
        default: {
            go sortItfin(sli, ch, &wg)
            wg.Add(1)
        }
        }
        wg.Wait()
        count++
    }   
    appendIt(howl, ch)
}

// Generates a slice of integers. Gathered from user as string, then converted
// into int and added to slice
func generateSlice() []int  {
    var ask string
    var slicey []int
    for (ask != "X") {
        fmt.Printf ("\n Enter the Integer to be added to the Slice ('X' to exit) ")
        fmt.Scan(&ask)     
        // check for exit condition
        if ask != "X" {
            i, _ := strconv.Atoi(ask)
            slicey = append(slicey, i)
        } 
    }
    fmt.Println("\n--- The Unsorted array------> ", slicey)
    return slicey
}


func sortIt1(items []int, ch chan int) {
    fmt.Println("\n--- Sorting First Segment--->", items)
    sort.Ints(items)
    fmt.Println("--- Sorted First Segment---->", items, "\n")
    for i := range items {
        ch <- items[i]
    }
} 

func sortIt2(items []int, ch chan int) {
    fmt.Println("\n--- Sorting Second Segment--->", items)
    sort.Ints(items)
    fmt.Println("--- Sorted Second Segment---->", items, "\n")
    for i := range items {
        ch <- items[i]
    }
} 

func sortIt3(items []int, ch chan int) {
    fmt.Println("\n--- Sorting Third Segment--->", items)
    sort.Ints(items)
    fmt.Println("--- Sorted Third Segment---->", items, "\n")
    for i := range items {
        ch <- items[i]
    }
} 

func sortItfin(items []int, ch chan int, wg *sync.WaitGroup) {
    fmt.Println("\n--- Sorting Final Segment--->", items)
    sort.Ints(items)
    fmt.Println("--- Sorted Final Segment---->", items, "\n")
    for i := range items {
        ch <- items[i]
    }
    wg.Done()
} 


// This go routine takes from the channel and puts integers into the final slice.
func appendIt(howl int, ch chan int) {
    items := make([]int, 0, howl)
    for i := 0; i < howl; i++ {
        temp := <- ch
        items = append(items, temp)
    } 
    sort.Ints(items)
    fmt.Println("\n--- The Sorted Array is: ", items)
}
