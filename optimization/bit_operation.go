// 交换两个数
func Swap(a, b *int) {
    if *a^*b != 0 {
	*a ^= *b
	*b ^= *a
	*a ^= *b
    }
}

// 交换符号
func ExcangeSymbol(a int) (ret int) {
    ret = ^a + 1
    return
}

// 求绝对值 
func CalcAbs(a int) (ret int) {
    ret = (a ^ a>>31) - a>>31
    return
}

// 二进制逆序 
func ByteReverse(a int) int {
    a = ((a & 0xAAAA) >> 1) | ((a & 0x5555) << 1)
    a = ((a & 0xCCCC) >> 2) | ((a & 0x3333) << 2)
    a = ((a & 0xF0F0) >> 4) | ((a & 0x0F0F) << 4)
    a = ((a & 0xFF00) >> 8) | ((a & 0x00FF) << 8)

    return a
}

// 求1的个数
func CalcOneNum(a int) int {
    a = ((a & 0xAAAA) >> 1) + (a & 0x5555)
    a = ((a & 0xCCCC) >> 2) + (a & 0x3333)
    a = ((a & 0xF0F0) >> 4) + (a & 0x0F0F)
    a = ((a & 0xFF00) >> 8) + (a & 0x00FF)

    return a
}
