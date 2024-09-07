package iteration

func Repeat(name string, num int) string {
	var res string

	for i := 0; i < num; i++ {
		res += name
	}

	return res
}