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
	if slc == nil {
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

// 二分法向前查找目标值位置
func (this optSliceInt) FindAcs(dest int) (index int) {
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
func (this optSliceInt) FindDesc(dest int) (index int) {
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
func (this optSliceInt) Index(dest int) (index int) {
	for i := 0; i < len(this.slc); i++ {
		if this.slc[i] == dest {
			return i
		}
	}
	return -1
}

// 查找目标值最后一次出现的位置
func (this optSliceInt) LastIndex(dest int) (lastIndex int) {
	for i := len(this.slc) - 1; i >= 0; i-- {
		if this.slc[i] == dest {
			return i
		}
	}
	return -1
}
