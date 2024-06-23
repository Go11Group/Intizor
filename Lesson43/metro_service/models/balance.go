package models

type Balance struct {
	UserId string 
	Amount float64 
}

type BalanceResp struct {
	Balance      float64 
	BalanceStatus string 
}