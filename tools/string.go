package tools

import (
	"sort"
	"strconv"
)

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func StringGroupToIntGroup(str []string) (intGroup []int) {
	for _, v := range str {
		i, _ := strconv.Atoi(v)
		intGroup = append(intGroup, i)
	}
	return
}

func StrContains(target string, strArr []string) bool {
	sort.Strings(strArr)
	index := sort.SearchStrings(strArr, target)
	if index < len(strArr) && strArr[index] == target {
		return true
	}
	return false
}
