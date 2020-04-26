package gslice

import "strconv"

// ============================================================================================================================================
//														string slice opt
// ============================================================================================================================================

// string 切片操作类
type optSliceString struct {
	slc []string
}

// 新建string 处理类
func NewOptSliStr(slc []string) optSliceString {
	if len(slc) == 0 {
		slc = []string{}
	}
	return optSliceString{slc}
}

// slice(string类型)查包含
func (this optSliceString) In(t string) bool {
	for i := 0; i < len(this.slc); i++ {
		if t == this.slc[i] {
			return true
		}
	}
	return false
}

// slice(string类型)元素去重
func (this optSliceString) Unique() []string {
	result := make([]string, 0)
	tempMap := map[string]struct{}{}
	ok := false
	for i := 0; i < len(this.slc); i++ {
		if _, ok = tempMap[this.slc[i]]; !ok {
			tempMap[this.slc[i]] = struct{}{}
			result = append(result, this.slc[i])
		}
	}
	return result
}

// 转换 string 切片为 int 切片
func (this optSliceString) ToSliInt() (res []int) {
	for i := 0; i < len(this.slc); i++ {
		if iValue, err := strconv.Atoi(this.slc[i]); err != nil {
			panic(err)
		} else {
			res = append(res, iValue)
		}
	}
	return
}

// 转换 string 切片为 int64 切片
func (this optSliceString) ToSliI64() (res []int64) {
	for i := 0; i < len(this.slc); i++ {
		if i64Value, err := strconv.ParseInt(this.slc[i], 10, 64); err != nil {
			panic(err)
		} else {
			res = append(res, i64Value)
		}
	}
	return
}

// string 类型切片转为 空接口切片
func (this optSliceString) SliTransIF() (resSli []interface{}) {
	for i := 0; i < len(this.slc); i++ {
		resSli = append(resSli, this.slc[i])
	}
	return
}
