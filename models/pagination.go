package models

// 分页查找结果
type Pagination struct {
	Page  int         `json:"page" desc:"第几页"`
	Size  int         `json:"size" desc:"每页几条记录"`
	Nums  int         `json:"nums" desc:"一共多少条记录"`
	Pages int         `json:"pages" desc:"总共几页"`
	Data  interface{} `json:"data" desc:"当前页的所有记录"`
}
