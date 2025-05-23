package Base

import (
	Interface "Bot/Interface"
)

type Button struct {
	Name Interface.ITemplate
	Data string
}

func (b *Button) GetTemplate() Interface.ITemplate {
	return b.Name
}

func (b *Button) GetData() string {
	return b.Data
}

func NewSimpleButton(data string) Interface.IButton {
	return &Button{
		Name: Template{Code: data, Text: data},
		Data: data,
	}
}
