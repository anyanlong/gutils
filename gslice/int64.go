package gslice

import "strconv"

// ============================================================================================================================================
//														int64 slice opt
// ============================================================================================================================================
// 判断函数定义
var (
	// 判断 参数1 是否 小于 参数2 的函数
	lessI64F = func(s, t int64) bool { return s < t }

	// 判断 参数1 是否 小于等于 参数2 的函数
	lessEqualI64F = func(s, t int64) bool { return s <= t }
)

// int64 切片操作类
type optSliceInt64 struct {
	slc []int64
}

// 新建int64 处理类
func NewOptSliI64(slc []int64) optSliceInt64 {
	if slc == nil {
		slc = []int64{}
	}
	return optSliceInt64{slc}
}

// int64切片的包含操作
func (this optSliceInt64) In(t int64) bool {
	for i := 0; i < len(this.slc); i++ {
		if t == this.slc[i] {
			return true
		}
	}
	return false
}

// int64切片的去重
func (this optSliceInt64) Unique() (resSlc []int64) {
	resSlc = make([]int64, 0)
	tempMap := map[int64]struct{}{}
	ok := false
	for i := 0; i < len(this.slc); i++ {
		if _, ok = tempMap[this.slc[i]]; !ok {
			tempMap[this.slc[i]] = struct{}{}
			resSlc = append(resSlc, this.slc[i])
		}
	}
	return resSlc
}

// slice(int64类型) 判断是否有列表值 小于目标值
func (this optSliceInt64) OrLess(t int64) bool {
	return this.orOpt(lessI64F, t)
}

// slice(int64类型) 判断是否所有有列表值 小于目标值
func (this optSliceInt64) AndLess(t int64) bool {
	return this.andOpt(lessI64F, t)
}

// slice(int64类型) 判断是否有列表值 小于目标值
func (this optSliceInt64) OrLessEqual(t int64) bool {
	return this.orOpt(lessEqualI64F, t)
}

// slice(int64类型) 判断是否有列表值 小于目标值
func (this optSliceInt64) AndLessEqual(t int64) bool {
	return this.andOpt(lessEqualI64F, t)
}

// 只要有一个满足 f函数 就返回true
func (this optSliceInt64) orOpt(f func(s, t int64) bool, t int64) bool {
	for i := 0; i < len(this.slc); i++ {
		if f(this.slc[i], t) {
			return true
		}
	}
	return false
}

// 只要全部满足 f函数 才返回true
func (this optSliceInt64) andOpt(f func(s, t int64) bool, t int64) bool {
	for i := 0; i < len(this.slc); i++ {
		if !f(this.slc[i], t) {
			return false
		}
	}
	return len(this.slc) > 0
}

// 转换int64切片为string
func (this optSliceInt64) ToSliString() (res []string) {
	for i := 0; i < len(this.slc); i++ {
		res = append(res, strconv.FormatInt(this.slc[i], 10))
	}
	return
}

// 转换 int64 切片为 int 切片
func (this optSliceInt64) ToSliInt() (res []int) {
	for i := 0; i < len(this.slc); i++ {
		res = append(res, int(this.slc[i]))
	}
	return
}

// int64类型切片转为 空接口切片
func (this optSliceInt64) SliTransIF() (resSli []interface{}) {
	for i := 0; i < len(this.slc); i++ {
		resSli = append(resSli, this.slc[i])
	}
	return
}
