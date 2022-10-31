package mock_mysql

func checknow(nowType int, max int) (next int, now int) {
	now = nowType
	if next = nowType + 1; next >= max {
		next = nowType
	}
	return
}
