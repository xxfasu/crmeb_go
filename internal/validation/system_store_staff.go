package validation

type GetSystemStoreStaffList struct {
	StoreID int64 `json:"storeId"`
	PageParam
}
