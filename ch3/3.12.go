package main

func main() {
	println(compare("我爱瑶瑶", "瑶瑶爱我"))
}

func compare(s string, s1 string) bool {
	if len(s) != len(s1) {
		return false
	}
	var _map = map[string]int{}
	var _map1 = map[string]int{}
	for _, v := range s {
		_map[string(v)]++
	}
	for _, v := range s1 {
		_map1[string(v)]++
	}
	for k, v := range _map {
		v1, ok := _map1[k]
		if !ok || v1 != v {
			return false
		}
	}
	return true
}
