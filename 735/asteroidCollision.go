package asteroidCollision

func asteroidCollision(asteroids []int) []int {
	res := make([]int, 0)
	for _, v := range asteroids {
		ok := true
		for ok && len(res) > 0 && res[len(res)-1] > 0 && v < 0 {
			pre := res[len(res)-1]
			if pre <= -v {
				res = res[:len(res)-1]
			}
			if pre >= -v {
				ok = false
			}
		}
		if ok {
			res = append(res, v)
		}
	}
	return res
}

//-1 1
//-1 -1
//1 1
//1 -1
