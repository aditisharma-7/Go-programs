package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func sortarr(arr []int32) []int32 {
    var temp int32
    for j := 0; j < len(arr); j++ {
        for i := 0; i < len(arr)-1; i++ {
            if arr[i] > arr[i+1] {
                temp = arr[i]
                arr[i] = arr[i+1]
                arr[i+1] = temp
            }
        }
    }
    //fmt.Println(arr)
    return arr
}

func miniMaxSum(arr []int32) {
    var sum1 int64
    var sum2 int64
    for i := 0; i < 4; i++ {
        sum1 += int64(arr[i])
    }
    for i := 1; i < 5; i++ {
        sum2 += int64(arr[i])
    }
    fmt.Print(sum1, sum2)
    //fmt.Println(sum2)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < 5; i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }
    arr = sortarr(arr)
    miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
