package Base

import (
	Interface "Bot/Interface"
)

type Template struct {
	Code string
	Text string
	Args []string
}

func (t Template) GetArgs() []string { return t.Args }

func (t Template) IsTranslated() bool { return t.Code == "" }

func (t Template) GetTemplateText() string { return t.Text }

func (t Template) GetTemplateCode() string { return t.Code }

// Hard code template
const (
	OnBackToParent string = "onBackToParent"
	OnCansel       string = "onCansel"
)

var OnBackToParentTemplate Interface.ITemplate = Template{Code: OnBackToParent, Text: "Back"}
var OnNextListPageTemplate Interface.ITemplate = Template{Code: OnNextListPage, Text: "Next Page"}
var OnPrefListPageTemplate Interface.ITemplate = Template{Code: OnPrefListPage, Text: "Pref Page"}
var OnCanselTemplate Interface.ITemplate = Template{Code: OnCansel, Text: "Cansel"}
