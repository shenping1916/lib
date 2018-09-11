// round 到最近的2的倍数
func minBuckets(v int) int {
    v--
    v |= v >> 1
    v |= v >> 2
    v |= v >> 4
    v |= v >> 8
    v |= v >> 16
    v++

    return v
}
