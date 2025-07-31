package models

type Dashboard struct {
	TotalIncome   string `json:"totalincome"`
	TotalExpenses string `json:"totalexpenses"`
}

// use depends on what router we decide to go with(mux or gin)
