var gBacktraceMutex sync.Mutex

func printBacktrace(err interface{}) {
    gBacktraceMutex.Lock()
    defer gBacktraceMutex.Unlock()
    fmt.Printf("panic: %v\n", err)

    i := 2
    for {
        pc, file, line, ok := runtime.Caller(i)
        if !ok {
            break
        }
        f := runtime.FuncForPC(pc)
        fmt.Printf("%d(%s): %s:%d\n", i-1, f.Name(), file, line)
        i++
    }
    fmt.Println("")
}
