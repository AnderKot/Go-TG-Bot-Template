package Base

import (
	Interface "Bot/Interface"
)

// Keyboard >>
type Keyboard struct {
	Rows []Interface.IKeyRow
}

func (k *Keyboard) GetRows() []Interface.IKeyRow {
	return k.Rows
}

// Keyboard <<

// KeyRow >>
type KeyRow struct {
	Keys []Interface.IKey
}

func (k *KeyRow) GetKeys() []Interface.IKey {
	return k.Keys
}

// KeyRow <<

// Key >>
type Key struct {
	Name Interface.ITemplate
	Data string
}

func (k Key) GetTemplate() Interface.ITemplate {
	return k.Name
}

func (k Key) GetData() string {
	return k.Data
}

// Key <<
