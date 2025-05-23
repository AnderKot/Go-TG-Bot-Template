package Controller

import (
	Database "Bot/Database"
	Table "Bot/Database/Repository/Table"
	Interface "Bot/Interface"
	Base "Bot/UI/Base"
	"strconv"
)

type ProfileState string

const (
	main         ProfileState = "main"
	editNickName ProfileState = "editNickName"
	editSteamID  ProfileState = "editSteamID"
)

type Profile struct {
	Database Database.Database
	ID       int64
	profile  Table.Profile

	state        ProfileState
	isStateCheng bool
	isNeedToPar  bool
}

func (m *Profile) Init() bool {
	m.state = main
	var err error
	m.profile, err = m.Database.Profile.GetProfile(m.ID)
	if err != nil {
		m.profile.NickName = "user" + strconv.FormatInt(m.ID, 10)
		return false
	}
	return true
}

func (m *Profile) OnProcessingMessage(text string) {
	switch m.state {
	case editNickName:
		{
			m.profile.NickName = text
			m.Database.Profile.UpdateProfile(m.profile)
		}
	case editSteamID:
		{
			m.profile.SteamID = text
			m.profile.IsValidSteamID = false
			m.Database.Profile.UpdateProfile(m.profile)
		}
	default:
		return
	}
}

func (m *Profile) GetButtons() []Interface.IButton {
	mi := make([]Interface.IButton, 0)

	switch m.state {
	case editNickName, editSteamID:
		{
			mi = append(mi, Base.NewSimpleButton(Base.OnCansel))
		}
	default:
		mi = append(mi, Base.NewSimpleButton(string(editNickName)))
		mi = append(mi, Base.NewSimpleButton(string(editSteamID)))
	}

	return mi
}

func (m *Profile) OnProcessingKey(data string) {
	switch m.state {
	case main:
		{
			switch data {
			case string(editSteamID):
				{
					m.state = editSteamID
				}
			case string(editNickName):
				{
					m.state = editNickName
				}
			}
		}
	case editNickName, editSteamID:
		{
			if data == Base.OnCansel {
				m.state = main
			}
		}
	}

}

func (m *Profile) GetTemplate() Interface.ITemplate {
	isValidSteamID := "-"
	if m.profile.IsValidSteamID {
		isValidSteamID = "+"
	}

	return Base.Template{
		Code: "ProfileMenu",
		Text: "[Profile Menu]" +
			"Nickname %s" +
			"Steam ID %s valid %s",
		Args: []string{
			m.profile.NickName,
			m.profile.SteamID,
			isValidSteamID,
		},
	}
}

func (m *Profile) GetNextPage() Interface.IConstructor {
	return nil
}

func (m *Profile) IsStageChanged() bool {
	return m.isStateCheng
}

func (m *Profile) IsNeedToParent() bool {
	return m.isNeedToPar
}

func (m *Profile) Finish() {
	m.Database.Profile.UpdateProfile(m.profile)
}
