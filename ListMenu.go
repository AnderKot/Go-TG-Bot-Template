package main

import (
	"fmt"
	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"slices"
	"strconv"
	"strings"
)

func ListMenu(stack CallStack) CallStack {
	if stack.Data == nil {
		return ReturnOnParent(stack)
	}
	stack.Action = ListMenu
	data := userDatas[stack.ChatID]

	mode := stack.Data["mode"]
	var items []string
	var page int
	if stack.IsPrint {
		stack.IsPrint = false
		// Print UI
		msg := tgBotAPI.NewMessage(stack.ChatID, GetTemplate(mode, data.languageСode))

		var keyboard tgBotAPI.InlineKeyboardMarkup

		pageStr, isPage := stack.Data["page"]
		if isPage {
			page, _ = strconv.Atoi(pageStr)
			items = GetItems(mode, page)
			keyboard = CreateInlineKaybord(3, items)

			if NextExist(mode, page) {
				keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, tgBotAPI.NewInlineKeyboardRow(
					tgBotAPI.NewInlineKeyboardButtonData(GetTemplate("back", data.languageСode), "back"),
					tgBotAPI.NewInlineKeyboardButtonData(GetTemplate("next", data.languageСode), "next"),
				))
			} else {
				keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, tgBotAPI.NewInlineKeyboardRow(
					tgBotAPI.NewInlineKeyboardButtonData(GetTemplate("back", data.languageСode), "back"),
				))
			}
		}

		finStr, isFind := stack.Data["find"]
		if isFind {
			items = FindItems(finStr)

			if len(items) == 0 {
				msg.Text += fmt.Sprintf(GetTemplate("noFound", data.languageСode), finStr)
			} else {
				keyboard = CreateInlineKaybord(3, items)
				keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, tgBotAPI.NewInlineKeyboardRow(
					tgBotAPI.NewInlineKeyboardButtonData(GetTemplate("back", data.languageСode), "back"),
				))
			}
		}

		msg.ReplyMarkup = keyboard
		_, _ = stack.Bot.Send(msg)

		mainMenuKeyboard := tgBotAPI.NewRemoveKeyboard(true)
		msg = tgBotAPI.NewMessage(stack.ChatID, "")
		msg.ReplyMarkup = mainMenuKeyboard
		_, _ = stack.Bot.Send(msg)
	}

	if stack.Update != nil {
		if stack.Update.Message != nil {
			_, isFind := stack.Data["find"]
			if isFind {
				stack.IsPrint = true
				stack.Data["find"] = stack.Update.Message.Text
				stack.Update = nil
				return ListMenu(stack)
			}

			return ListMenu(CallStack{
				ChatID:  stack.ChatID,
				Bot:     stack.Bot,
				IsPrint: true,
				Parent:  &stack,
				Data: map[string]string{
					"mode": stack.Data["mode"],
					"find": stack.Update.Message.Text,
				}})
		}

		if stack.Update.CallbackQuery != nil {
			choice := stack.Update.CallbackQuery.Data
			switch choice {
			case "back":
				{
					return ReturnOnParent(stack)
				}
			case "next":
				{
					stack.Update = nil
					stack.IsPrint = true
					return ListMenu(CallStack{
						ChatID:  stack.ChatID,
						Bot:     stack.Bot,
						IsPrint: true,
						Parent:  &stack,
						Data: map[string]string{
							"mode": stack.Data["mode"],
							"page": strconv.Itoa(page + 1),
						}})
				}
			default:
				if ItemExist(choice) {
					stack.Update = nil
					stack.IsPrint = true
					// Here it is worth placing the choice of the Run to be launched depending on the operating mode of the list
					if stack.Data["mode"] == "RunTemplate" {
						return RunTemplate(CallStack{
							ChatID:  stack.ChatID,
							Bot:     stack.Bot,
							IsPrint: true,
							Parent:  &stack,
							Data: map[string]string{
								"item": choice,
							}})
					} else {
						return Chop(stack)
					}
				}
			}
		}
	}

	return stack
}


func CreateInlineKaybord(countInRow int, items []string) tgBotAPI.InlineKeyboardMarkup {
	inlineKeyboard := tgBotAPI.NewInlineKeyboardMarkup()
	var row []tgBotAPI.InlineKeyboardButton

	for _, item := range items {
		row = append(row, tgBotAPI.NewInlineKeyboardButtonData(item, item))
		if len(row) == countInRow {
			inlineKeyboard.InlineKeyboard = append(inlineKeyboard.InlineKeyboard, row)
			row = []tgBotAPI.InlineKeyboardButton{}
		}
	}

	if len(row) > 0 {
		inlineKeyboard.InlineKeyboard = append(inlineKeyboard.InlineKeyboard, row)
	}

	return inlineKeyboard
}


// Just for exemple
var AllItems = []string{
	"item 1",
	"item 2",
	"item 3",
	"item 4",
	"item 5",
	"item 6",
	"item 7",
	"item 8",
	"item 9",
	"item 10",
	"item 11",
	"item 12",
	"item 13",
	"item 14",
	"item 15",
}

func NextExist(mode string, page int) bool {
	return len(AllItems) > 10+(9*page) // Just for exemple
}

func ItemExist(name string) bool {
	return slices.Contains(AllItems, name) // Just for exemple
}

func FindItems(name string) []string {
	// Just for exemple >>
	var items []string

	for _, item := range AllItems {
		if strings.Contains(item, name) {
			items = append(items, item)
		}
	}

	return items
	// Just for exemple <<
}

func GetItems(mode string, page int) []string {
	// Just for exemple >>
	requestLastIndex := 9 + (9 * page)
	maxIndex := len(AllItems)
	if maxIndex < requestLastIndex {
		requestLastIndex = maxIndex
	}

	return AllItems[(9 * page):requestLastIndex]
	// Just for exemple <<
}
