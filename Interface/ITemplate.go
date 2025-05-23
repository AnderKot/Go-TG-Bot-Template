package Interface

type ITemplate interface {
	IsTranslated() bool
	GetTemplateCode() string
	GetTemplateText() string
	GetArgs() []string
}
