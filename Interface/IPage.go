package Interface

type IPage interface {
	// Common
	GetName() string

	// Input
	OnProcessingMessage(text string)
	OnProcessingKey(keyData string)

	// Navigation
	OnGetNextPage() IConstructor
	OnBackToParent() bool

	// Print
	GetMessageText() ITemplate
	GetKeyboard() IKeyboard
}
