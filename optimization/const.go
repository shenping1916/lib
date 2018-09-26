const (
    // 最大uint
    MaxUint     = ^uint(0)                    // 1<<32 - 1 or 1<<64 - 1

    // 最大int
    MaxInt      = int(MaxUint >> 1)           // 1<<31 - 1 or 1<<63 - 1

    // 最小int
    MinInt      = int(-MaxInt - 1)            // -(1<<31) or -(1<<63)

    // 计算机位数：32 / 64
    BitsPerWord = int(32 * (1 + MaxUint>>63)) // 32 or 64
)
