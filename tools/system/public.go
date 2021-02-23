package system

type Child struct {
	Id  int
	Pid int
}

func HasChild(arr []Child, id int) bool {
	for _, v := range arr {
		if v.Pid == id {
			return true
		}
	}
	return false
}
