var (
	_strSet = make(map[string]struct{}, 1<<8)
	_arr    = make([]string, 0, 4)
	_opt    = make([]string, 0, 2)
)

func init() {
	for i := 0; i < 1<<8; i++ {
		_strSet[strconv.Itoa(i)] = struct{}{}
	}
}

func dfs(s string, st int) {
	defer func() {
		if len(_arr) > 0 {
			_arr = _arr[:len(_arr)-1]
		}
	}()

	if len(_arr) == 4 {
		if st == len(s) {
			_opt = append(_opt, strings.Join(_arr, "."))
		}
		return
	}

	for offset := 1; st+offset <= len(s); offset++ {
		if _, ok := _strSet[s[st:st+offset]]; !ok {
			break
		}

		_arr = append(_arr, s[st:st+offset])
		dfs(s, st+offset)
	}
}

func restoreIpAddresses(s string) []string {
	_arr = _arr[:0]
	_opt = _opt[:0]
	dfs(s, 0)
	return _opt
}