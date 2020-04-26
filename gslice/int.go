package gslice

import "strconv"

// ============================================================================================================================================
//														int slice opt
// ============================================================================================================================================
// int 切片操作类
type optSliceInt struct {
	slc []int
}

// 新建int 处理类
func NewOptSliInt(slc []int) optSliceInt {
	if len(slc) == 0 {
		slc = []int{}
	}
	return optSliceInt{slc}
}

// slice(int类型)查包含
func (this optSliceInt) In(t int) bool {
	for i := 0; i < len(this.slc); i++ {
		if t == this.slc[i] {
			return true
		}
	}
	return false
}

//  slice(int类型)元素去重
func (this optSliceInt) Unique() []int {
	result := make([]int, 0)
	tempMap := map[int]struct{}{}
	ok := false
	for i := 0; i < len(this.slc); i++ {
		if _, ok = tempMap[this.slc[i]]; !ok {
			tempMap[this.slc[i]] = struct{}{}
			result = append(result, this.slc[i])
		}
	}
	return result
}

// 转换 int 切片为 string 切片
func (this optSliceInt) ToSliString() (res []string) {
	for i := 0; i < len(this.slc); i++ {
		res = append(res, strconv.Itoa(this.slc[i]))
	}
	return
}

// 转换 int 切片为 int64 切片
func (this optSliceInt) ToSliI64() (res []int64) {
	for i := 0; i < len(this.slc); i++ {
		res = append(res, int64(this.slc[i]))
	}
	return
}

// int类型切片转为 空接口切片
func (this optSliceInt) SliTransIF() (resSli []interface{}) {
	for i := 0; i < len(this.slc); i++ {
		resSli = append(resSli, this.slc[i])
	}
	return
}
