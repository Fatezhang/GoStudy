package dataType

type man struct {
	name string
	age  int8
	set  bool
}

func GetMan() man {
	return man{"jiaheng", 23, true}
}
