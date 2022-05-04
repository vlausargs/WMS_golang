package db_data

type TransferTxParams struct {
	FromUserID int64 `json:"from_user_id"`
	ToUserID   int64 `json:"to_user_id"`
	Amount     int64 `json:"amount"`
}

// type TransferTxResult struct {
// 	FromUser db.User `json:"from_user"`
// 	ToUser   db.User `json:"to_user"`
// }
