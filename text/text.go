package text

type Text struct {
	Ja string
	En string
}

func Default() Text {
	texts := []Text{
		{"動けばヨシ！", "If it moves, it's okay."},
		{"ヨシ！", "Great!"},
		{"LGTM！", "LGTM!"},
		{"一日は24時間ある", "One day has 24 hours"},
		{"今日中という意味は明日の朝まで", "The meaning of today is until tomorrow morning"},
		{"祈れば動く。", "Move if you pray."},
		{"技術的には可能です。", "It is technically possible."},
		{"変態！", "transformation!"},
		{"ヤバい！", "It was awesome"},
	}

	return texts[RandomNumber(len(texts))]
}

func Motto() Text {
	texts := []Text{
		{"天才か？", "Are you a genius?"},
		{"天使だったか", "You're an angel."},
		{"救世主", "You're a savior."},
		{"こいつ･･･できる！", "You can do anything."},
		{"こいつ･･･狂ってる！", "You're crazy!"},
	}

	return texts[RandomNumber(len(texts))]
}
