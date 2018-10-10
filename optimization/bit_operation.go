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

// int转字符串,wid为int宽度
// 转换后append到[]byte
// 类似c的iota写法
// 来自官方标准库：/src/log/log.go
func itoa(buf *[]byte, i int, wid int) {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}

	// i < 10
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}

// 求1的个数
func CalcOneNum(a int) int {
    a = ((a & 0xAAAA) >> 1) + (a & 0x5555)
    a = ((a & 0xCCCC) >> 2) + (a & 0x3333)
    a = ((a & 0xF0F0) >> 4) + (a & 0x0F0F)
    a = ((a & 0xFF00) >> 8) + (a & 0x00FF)

    return a
}
