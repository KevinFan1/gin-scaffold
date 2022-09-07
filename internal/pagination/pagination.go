package pagination

import (
	"code/gin-scaffold/schemas"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
)

type Params struct {
	DB      *gorm.DB
	Page    int
	Size    int
	OrderBy []string
	ShowSQL bool
}

type Paginator struct {
	TotalRecord int64 `json:"total"`
	TotalPage   int   `json:"total_page"`
	Items       any   `json:"items"`
	CurrentPage int   `json:"current_page"`
	PrePage     int   `json:"pre_page"`
	HasPre      bool  `json:"has_pre"`
	NextPage    int   `json:"next_page"`
	HasNext     bool  `json:"has_next"`
}

func Paging(p *Params, resultVal any) (*Paginator, error) {
	maxSize := 20
	db := p.DB

	if p.ShowSQL {
		db = db.Debug()
		fmt.Println("Debug...")
	}

	// 初始化page size
	if p.Page < 1 {
		p.Page = 1
	}

	if p.Size < 1 || p.Size > maxSize {
		p.Size = maxSize
	}

	// 排序
	if len(p.OrderBy) > 0 {
		for _, field := range p.OrderBy {
			db = db.Order(field)
		}
	}

	done := make(chan bool, 1)
	var count int64
	var offset int
	go getCount(db, resultVal, done, &count)
	<-done

	paginator := &Paginator{
		TotalPage:   int(math.Ceil(float64(count) / float64(p.Size))),
		TotalRecord: count,
		Items:       resultVal,
		CurrentPage: p.Page,
		HasPre:      false,
		HasNext:     false,
	}

	if p.Page > paginator.TotalPage {
		return nil, errors.New("查询列表失败")
	}

	offset = (p.Page - 1) * p.Size

	db.Limit(p.Size).Offset(offset).Find(resultVal)
	// 判断是否有前一页
	if p.Page > 1 {
		paginator.PrePage = p.Page - 1
		paginator.HasPre = true
	} else {
		paginator.PrePage = p.Page
	}

	// 判断是否有后一页
	if p.Page < paginator.TotalPage {
		paginator.NextPage = p.Page + 1
		paginator.HasNext = true
	} else {
		paginator.NextPage = p.Page
		paginator.HasNext = false
	}

	return paginator, nil

}

func getCount(db *gorm.DB, T any, done chan bool, count *int64) {
	db.Model(T).Count(count)
	done <- true
}

func Scan[T any](c *gin.Context, db *gorm.DB, result []T) (paginator *Paginator, err error) {
	var params schemas.QueryPaginatorParams
	err = c.ShouldBindQuery(&params)
	paginator, err = Paging(&Params{
		DB:      db,
		Page:    params.Page,
		Size:    params.Size,
		OrderBy: []string{},
		ShowSQL: false,
	}, &result)
	return
}
