package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "3.3" )

func a7euq5Q6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wUNwoDtrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dl0W5lQFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gV9zfzUbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i3Fi0Q5pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cW2mWPWOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B6tn1eIPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kTFHnxh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qRuzFWJsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ztt4hbg3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yKWE6Dl3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F59kKhnNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dFK4AuBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KeKgC1iGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u8bMw3oTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vZmi2kr0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1GltlSp2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gbYezwzcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cGyIJeJKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8Up6gw16Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u5K2BXZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TsypoJ8gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XpKrZWo0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fYnppYbZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PfkCRXZwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XP0IoiVsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xODlMiPSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gEAI3hjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mnLm7cBzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Z7DGmrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hn3NMlXgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JhbZiuQYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OBhCOvWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4uICmJIuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UAT9MceJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vanBd4q9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dokmEToXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IANtnRa0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bSfVOiAuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func POYzXaDIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 89gc63bwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MS69hKXXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EIUH0BvQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3ybnLTJiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MnKvbzMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PFUkqakoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ulZPHo5yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VbIrDlIGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p8eRSS75Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y1AIWfahWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MNiEUMVbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jOUMrduhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D8pay1J8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 04b0ESjtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sQeAgAU6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AbmPBaBIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bOnqDq8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3YFFP48XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MKwgH14fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 21awcfo3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jhz4ZdIjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func azKMgW8rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KWrPsxq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MaZIG8zhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yseZ4WSFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wr1x1bqpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GH35ydqRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wb4omrc2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5qKsS5DdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u0l7AHSBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sUy1DRS9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cfJ72u5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KKeU1rxkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xs1pZtkTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bAWgJmCxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JpXojNrMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZGZbWFfRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IhOqiK83Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G8kvqAqOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E94NUaZnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5Ml32KyxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zPOjZXlvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8QMEyrbzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 673qhFoPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5THKNMN2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 34Ku3eRHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 88esapbAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IsSGI1UEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t5eTI51kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fWbWAnqvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vnkaufRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aQBwzza9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9CQQZ5fYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j8nMwTQAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fJ1zC4XmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kljxN2aJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Io4733ocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FUqCtvKdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GDutud9rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JJqTwyrKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WIioY7hEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OYCZcGbRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j6VTJMJaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PebabFAlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 05SZ0bZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iI3Fc1tqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u93wqlCDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sP7MFMfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t6uSDpjhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dfDelGdmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 20kz0fKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8vHmoBFhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9IjeFaVYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oKFZfqTgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bZ1NKBmzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JxJpLx9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CS9MYZgLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LNGIQDyVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ut2z3TgPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kcbVXlX1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ByvlH2SfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6gvpCCAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BTLFV32sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eWkPu94gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p0Fg1c7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZYoO6IsnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xaBfIorAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SJ7ypufdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2tL2dpfnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hpWYF5mDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8lj6nil0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZzXUkUMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H8ejiR6lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ma8N2MapWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ytGcRGomWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TRXA2ZiIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9gRvabWMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 14itBVynWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qjU34PYCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func deaKZKcsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZYCx004yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xtb4TmzqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RlXDNyZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ShXU8qUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v5e0LIJWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pDGonuR1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xNYDhwC9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FcAGHO45Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tMvalvjFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j1lU4MwDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b1tnoaw8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ol9oIQ4dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AIOscK5mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0gtA98CQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TPWCHGAXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WUJIeITsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oLwQCZ00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O19CzWrAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C5fGKob5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MyKV6z5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D79WuBYhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v60SMeUKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1zqklkJ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r2Uxj6hnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gVEd7dQlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bVnrBEaIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TDcQGmBJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fAEhZX2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iFCZ0YIpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5MXBy9yhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BCv4veU2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VVqppPpOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mn7vJrydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LkDiKDj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nKWQgnc2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w0NlOmY9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HdiGOyXcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qf00dFpAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cqWZVuhLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yjJHrnPNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lVTnbKftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 30x4M5aoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MpvEtAsoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z6viewlKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func be3ck6jAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qT1OOBwxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZbK7jL7tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EThhxKfqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vczZhaknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 27DBAfZHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vIdlq8M0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ofPa4ZVgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YMZ4wKL0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rZEjTTaZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LyhFt2BSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cH8crZQDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F6Oy9CZUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AmoygMhNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1B74BNvFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JCvCCmbXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func apa69ibwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func te9SpsshWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func exqAIlLpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4prbke0GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XKKCknpFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Whwrr9t4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IDl2o1j0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ufqyN9TIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JpRHmjKoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3pqZyf0cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KyUO07ujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XzvtZ3zoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oONDqbQjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yix5MMkWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bDiXbnGdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8NwhBx2RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Au6ssNv5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7wZ8ympqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZYedijFFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 74EtXKZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KGReD5gHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c0Jns0PdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KVEu12K1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nTYQ6QCEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wW2Yn0TGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dliYj5aHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rdwq7ZCsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UjAQoLhGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KkGoEm84Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n0jhoq60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yumwr0rcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EqgooYiyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8ZwvbHSuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func alRXqSwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0VhkYks5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GCIebNnLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l9Pt22AdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ThX18QEwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lHF9byIgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6dvvpBwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VI3q4ZpfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XJXK1chqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 26UUhfoFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X4eJwvyFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LvwdwjgLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sz4uNXgxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q06NxFpaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AZoK1CCJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func taSYRrjIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func THRF4RxGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lx6l7uWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BVneButfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SvoEEfHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f4pl6Y4JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9Mri57gQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K1z0PmWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KMT6SbAOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QEgLxNtOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lHrT0ZiSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KoTOWamkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vB9ohhdnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sjwJoo1WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x7BjtLJbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tTMuFt2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HT5evvhOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JUqJFz2UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7dGCGtQIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QlURggUlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XEAYkvfKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5oi5SLMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i7MoUN12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0y27XhDAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RvsflrKzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VCQ3YvlQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zor5uborWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G7nvys1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pq6ykDtuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7knZSQdDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OftTBEX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KbkQhKIfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z0M4t3PhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ISBEyVcaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GhgpryoxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2KLNIjPYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GzR7LbmyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3CuUHXbsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cfP7pZUBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h782mmDTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vPCLLEUsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yLL5BtNJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3O34YsvWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z32yficOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5FJ2uo7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QaxEybDJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Np24OnCLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x7KPtOc6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r1gjk6uCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UcHazJnuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QnZsiVMwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gQt2FlwxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kgggoi8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A19DwnAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vJ6e7P8xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QVaGye38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rQAmFjgWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hvnwfZp9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GEMFmjeAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XRpSnZPLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CeLREZ4mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qUt8pVXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3lPsicrRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3wNy2t3qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vlYjQKQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DXTGR23oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uC8dRLBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OsOH1GGCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AVdDSLfdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QWy0P5ZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l5S1PhJaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C1OfQIh3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8Qp6qUemWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sy3oThMTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pb1YViCRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4HsfMqTdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FSZ05041Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func evys12DIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iUL3pNO9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func al6SAHiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 14YXZQvEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8N208bjkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FtWanXHNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wsc5zkS9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VtjDae1QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZIGoPGLoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QEY8g8C0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hCrIeOtgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d0p6F3lgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g3tmvm5kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vU4H83NMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Eisa8U4wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kqOxsFNTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YCsXP4PjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vHtqOMsYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kvwD92NfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z2zCzxskWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RwXQL9BRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XqPcNHGoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lU9sdyBEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HpfjsXEeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Txe286qhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Veh9aT8YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Yw6AzEZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4KfW4lycWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2f4rdzhwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 09k5CIMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cvrQbZHYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AvLSQLfiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ucWHKg6ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aw2TSSlqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WSfMIj6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2eRQqKFjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rv8mzOJiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OILph07fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3cPU2JUYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FT43meT3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4bv4yYtMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6FaOBjVuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o19VtBemWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ItPncywtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3ijWlT1hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s9VwRedtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bqZnffAFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KGiA4vlzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q8ufbwMIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wTFixhJLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QwMWuWf0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LGcC5LE1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HY4KkShBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o4cgq4T3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OCaiC6iSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2n0CYhiIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9PGfBL9OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zkXBDs6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u44xVZBoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IbPZs5ZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func piddWnk1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b7vPyUS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rh5bRoeYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gmblNFsDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8fUoAGzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YRH9TTLkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4qPwxIfyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b77B00JkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wlqi0PRhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gwuvdb3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EhGjxATKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v3aNnv1yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1PzuHecOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aMwZ7DDNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5XdZSm9PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VqVqZhK4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MbEcawBJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L56EZBTuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B2AifCrSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dcGDxoDTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gfJcZSviWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QZ4rRlpNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m9DSCWVKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ne7UF7BJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ucgfPVOmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rw8ZvB6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func anHUHGTBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qUXpaAQWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hMLePQBLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H5MKAWIFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TxunjEUEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lECk67y1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C5h9BVLBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NHhsaPivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2AJLvdf8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Viu3300wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XYmuypKgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ryXEWZanWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E1tVUAbCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S6El0rJlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jjIsrsXiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fDmtUiuZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6UEmwWzoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5zNYsqf7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NsSIlPZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nJ5WJQgJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func csmMLwOfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HqEiUrpiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sC7taJpyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qyCaYFliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 19lRztGaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func clljmFgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YOL54kuCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fPI77rUyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SqdEPXJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8h7oKckHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BDKvl9TrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pPDTaCnoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4f6PXFpCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kZJNxEc8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mHTu8iWtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IADrVLklWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JqXYXqayWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 54e5ZAWlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VJ2OVZ7TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mXhX40c9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HiEpJvHxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i6qbcCZKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GaYbyVi3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jIUjqZDUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func II0jvQV2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NTUKrgJHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L56OfEZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6AYvZ0XIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NSlIPhhNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OqTYXHE3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2fWPzpolWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mNyGVaz9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YSOukH3OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QetlVvxSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QZGSgWoEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iVS6x485Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rFB4ntTkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x60a8JcuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QIKLC6VuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DInBgNxQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uMSBDrWbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7jkuUvUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func plx8p7EpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ypav24EnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func scpbqmEAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vvfWpAYkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NirYwR3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aq1UH8l2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GYB6XyDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1JNp3z7YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YRZ39fkPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TFD4v7frWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f1qzN2UsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qMXpGEl4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f720zXGDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hZqnbupiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FFdWAnSjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5Tc9vkObWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func taXDSwLBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1KynFshxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GPtFH9lFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7oUncafnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nTtfMQgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6T4yb79aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gws9gtpWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XFkygoOJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gw2hwBzFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t5Poz9aJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u6HFla7IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aUW1S0MFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zG8ld7r3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tbfQS9wxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5iBdKV2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qb0z3Yf8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FmLg09swWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0COghcKtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vou6KdBXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vKKQz1O7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jtxi7pF3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bk0sZQQyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UzNk6s9KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FJyhCGGlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DUh43oCGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FryDky9YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2k1ozE8SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DVIzImlnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y8oTbN0rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qa9A1SIxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KnjF2DeeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YdtS7mmyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z7ltDs2YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 20PkBhmpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RFyK4axMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fxco3P9vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z8WtV86lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3Mki97LTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 56XnMrCkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ogYaVRVYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QxW6NaP4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dd7I8fXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IxxOtoEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PUC4vm9wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hFasAjx2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7rCOfKy3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e7AeoeZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func spV0HlN4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sK1gBterWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zvz7UsC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3vdnwBMWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ChXhCoowWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vnT9OyfhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rTCcdQxQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FlBa59IXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n31ru6RCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gx1bCdScWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eBuXBDosWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QSM9GcXlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EzV21zQnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HQ9edbItWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XhDdFaxPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cdvLxBuWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xpAxvr2pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pyk2qUVSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xtUCkCloWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3zPHd3fNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1XPKv5iDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T1A3gD31Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zxvfZV5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ypaMZm6kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r6I8zZ0tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cWdNHvW1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BK4GwzqRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4nyBRw4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IDXSIPi7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jYposXS1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QWTUVctmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0eyIZoAwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wpEPW3P3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VYzajGBAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D8qOejkyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BxDxTBqTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SZF2y585Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M8cavxV6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eMOHHMKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C8izqZzxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TOXpwDRiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5KBAzptRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hrIzsMvVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IRhKQTlkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4TX2KzbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5NupBbSyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LSA4m6WfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rya09PeFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YNQ1ZKAmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2hOGznrTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EZUvw1EhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s1kgE7pBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KFo95ZdxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ac1MgIytWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bcDZJ5cmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5LzZW9dyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t99QbDA1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KMTzNGOXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t8hRCuBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wNncHCirWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MOBQnku4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dkX75Dk8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8rAm8sgFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ov3hdJxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1ydwHkgVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gkQSbEI6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gc6I2j3xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

