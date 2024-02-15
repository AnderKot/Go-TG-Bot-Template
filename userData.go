package main

type userData struct {
	// Акаунт
	userName     string
	firstName    string
	lastName     string
	languageСode string
	// Статистика
	totalMatches  int
	winPercentage float32
	// Баланс
	walletBalanse  float32
	walletCurrency string
	// Рефералы
	referralsCount int
	referralsCash  float32
}
