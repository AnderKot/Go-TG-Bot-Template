package Interface

type IKeyboard interface {
	GetRows() []IKeyRow
}

type IKeyRow interface {
	GetKeys() []IKey
}

type IKey interface {
	GetTemplate() ITemplate
	GetData() string
}
