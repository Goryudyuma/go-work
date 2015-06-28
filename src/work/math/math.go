package math

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		{
			s += a[i]
		}
	}
	return s
}
