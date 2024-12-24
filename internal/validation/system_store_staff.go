package validation

type GetSystemStoreStaffList struct {
	StoreID int64 `query:"storeId"`
	PageParam
}
