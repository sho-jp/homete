package text

type Text struct {
	Ja string
	En string
}

func Default() Text {
	texts := []Text{
		{"最高", "Great!"},
		{"強す", "Strong!"},
		{"す!", "Strong!"},
		{"すぎ!", "Strong!"},
	}

	return Random(texts)
}

func Motto() Text {
	texts := []Text{
		{"ほんと最高!", "Very Great!"},
		{"ほんと強す", "Very Strong!"},
		{"ほんとす!", "Very Strong!"},
		{"ほんとすぎ!", "Very Strong!"},
	}

	return Random(texts)
}
