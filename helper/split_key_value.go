package helper

func SplitKeyValue(evt string) []string {
	for i := 0; i < len(evt); i++ {
		if evt[i] == '=' {
			return []string{evt[:i], evt[i+1:]}
		}
	}
	return nil
}
