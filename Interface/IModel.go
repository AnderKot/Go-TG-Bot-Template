package Interface

type IModel interface {
	Init() bool
	OnProcessingMessage(text string)
	OnProcessingKey(data string)
	GetTemplate() ITemplate
	GetButtons() []IButton
	GetNextPage() IConstructor
	IsStageChanged() bool
	IsNeedToParent() bool
	Finish()
}

type IButton interface {
	GetTemplate() ITemplate
	GetData() string
}
