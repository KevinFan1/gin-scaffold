package schemas

type QueryPaginatorParams struct {
	Page int `form:"page,default=1"`
	Size int `form:"size,default=20"`
}
