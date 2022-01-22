package validator

const (
	intSize = 32 << (^uint(0) >> 63)
	MaxInt  = 1<<(intSize-1) - 1
	MinInt  = -1 << (intSize - 1)
	MaxUint = 1<<intSize - 1
)
