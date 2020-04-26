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
	if slc == nil {
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

// 二分法向前查找目标值位置
func (this optSliceString) FindAcs(dest string) (index int) {
	lo, hi := 0, len(this.slc)-1
	for lo <= hi {
		m := (lo + hi) >> 1
		if this.slc[m] < dest {
			hi = m - 1
		} else if this.slc[m] > dest {
			lo = m + 1
		} else {
			return m
		}
	}
	return -1
}

// 二分法向后查找目标值位置
func (this optSliceString) FindDesc(dest string) (index int) {
	lo, hi := 0, len(this.slc)-1
	for lo <= hi {
		m := (lo + hi) >> 1
		if this.slc[m] < dest {
			lo = m + 1
		} else if this.slc[m] > dest {
			hi = m - 1
		} else {
			return m
		}
	}
	return -1
}

// 查找目标值第一次出现的位置
func (this optSliceString) Index(dest string) (index int) {
	for i := 0; i < len(this.slc); i++ {
		if this.slc[i] == dest {
			return i
		}
	}
	return -1
}

// 查找目标值最后一次出现的位置
func (this optSliceString) LastIndex(dest string) (lastIndex int) {
	for i := len(this.slc) - 1; i >= 0; i-- {
		if this.slc[i] == dest {
			return i
		}
	}
	return -1
}
