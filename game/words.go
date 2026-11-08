package game

var englishWords = map[int][]string{
	1: {"a", "i"},

	2: {
		"go", "if", "in", "is", "it", "my", "no", "on", "or",
		"to", "up", "us", "we",
	},

	3: {
		"and", "any", "but", "can", "for", "get", "had", "has", "how", "its", "may",
		"new", "not", "now", "our", "out", "say", "see", "set", "she", "sit", "sky",
		"the", "too", "try", "use", "was", "way", "who", "why", "you",
	},

	4: {
		"also", "back", "been", "call", "came", "come", "down", "each", "even", "find",
		"from", "give", "good", "have", "help", "here", "high", "just", "know", "like",
		"line", "long", "look", "make", "many", "more", "most", "much", "must", "name",
		"need", "next", "only", "over", "plan", "read", "real", "same", "show", "some",
		"such", "take", "tell", "than", "that", "them", "then", "they", "this", "time",
		"true", "turn", "type", "unit", "upon", "used", "very", "want", "well", "went",
		"were", "what", "when", "will", "with", "word", "work", "year", "your",
	},

	5: {
		"about", "after", "again", "below", "bring", "could", "every", "first", "found",
		"great", "house", "large", "learn", "never", "other", "place", "plant", "point",
		"right", "small", "sound", "spell", "still", "study", "their", "there", "these",
		"thing", "think", "three", "water", "where", "which", "world", "would", "write",
		"young",
	},

	6: {
		"around", "become", "before", "because", "between",
	},

	7: {
		"another", "between", "brought", "changed", "example", "general", "however",
		"include", "instead", "nothing", "offered", "outside", "present", "problem",
		"provide", "reading", "several", "special", "started", "through", "whether",
		"without", "working", "written",
	},

	8: {
		"actually", "anything", "becoming", "included", "learning", "possible",
		"provided", "question", "together",
	},
}

var russianWords = map[int][]string{
	1: {
		"а", "в", "и", "к", "о", "с", "у", "я",
	},

	2: {
		"он", "мы", "вы", "ты", "их", "им", "не", "ни",
		"на", "но", "да", "за", "от", "до", "по", "ко",
		"во", "об", "из", "уж", "ах", "ух",
	},

	3: {
		"это", "был", "все", "нет", "так", "тут", "там", "где",
		"кто", "что", "мой", "наш", "ваш", "сам", "дом", "мир",
		"год", "лес", "сад", "кот", "рот", "нос", "сон", "ход",
		"бок", "век",
	},

	4: {
		"вода", "дело", "день", "ночь", "рука", "нога", "глаз", "лицо",
		"тело", "душа", "сила", "труд", "путь", "друг", "враг", "мама",
		"папа", "отец", "хлеб", "стол", "окно", "море", "река", "гора",
		"поле", "снег", "лето", "зима", "язык", "игра",
	},

	5: {
		"время", "место", "книга", "город", "улица", "земля", "ветер",
		"огонь", "голос", "мысль", "число", "точка", "линия", "форма",
		"смысл", "текст", "жизнь", "народ", "закон", "право", "часть",
		"конец", "месяц", "вечер", "начало", "ответ", "выбор", "файл",
	},

	6: {
		"работа", "машина", "воздух", "солнце", "звезда", "страна",
		"музыка", "память", "грусть", "любовь", "забота", "теория",
		"задача", "ошибка", "строка", "символ", "сервер", "клиент",
		"данные", "система", "свобода",
	},

	7: {
		"природа", "решение", "история", "желание", "порядок", "понятие",
		"общение", "функция", "создать", "проверка", "объект", "запрос",
		"комната", "энергия", "человек", "участок", "размер", "картина",
	},

	8: {
		"компьютер", "алгоритм", "телефон", "документ", "обучение",
		"источник", "качество", "значение", "движение", "культура",
		"внимание", "программа", "развитие", "создание", "практика",
	},
}
