package schema

// PaginationParam 分页查询条件
type PaginationParam struct {
	Pagination bool `form:"-"`                                     // 是否使用分页查询
	OnlyCount  bool `form:"-"`                                     // 是否仅查询count
	Current    uint `form:"current,default=1"`                     // 当前页
	PageSize   uint `form:"pageSize,default=10" binding:"max=100"` // 页大小
}

// PaginationResult 分页查询结果
type PaginationResult struct {
	Total    int  `json:"total"`
	Current  uint `json:"current"`
	PageSize uint `json:"pageSize"`
}
