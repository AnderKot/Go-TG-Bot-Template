package main

// MessageTemplates[MessType][MessLanguageСode]Template
var MessageTemplates = map[string]map[string]string{
	"RunTemplate": {
		"ru":  "[Шаблон меню]\nШаблон текста\n%s",
		"eng": "[Template menu]\nTemplate text",
	},
	"back": {
		"ru":  "⬅️ Назад",
		"eng": "⬅️ Back",
	},
}

func SelectTemplate(messageType string, messageLanguageСode string) string {
	var template = MessageTemplates[messageType][messageLanguageСode]
	if template == "" {
		template = MessageTemplates[messageType]["eng"]
	}
	if template == "" {
		template = "Template not specified !\n"
	}
	return template
}
