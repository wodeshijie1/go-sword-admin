package bo

//order 排序规则
type order struct {
	Column string `json:"column"`
	Asc    bool   `json:"asc"`
}

type dept struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type job struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type role struct {
	ID        int    `json:"id"`
	Level     int    `json:"level"`
	Name      string `json:"name"`
	DataScope string `json:"dataScope"`
}

//paging 分页器所含字段 (公共父类)
type paging struct {
	Current          int      `json:"current"`
	CountID          int      `json:"count_id"`
	MaxLimit         int      `json:"maxLimit"`
	Page             int      `json:"page"`
	SearchCount      int      `json:"searchCount"`
	Size             int      `json:"size"`
	Total            int      `json:"total"`
	HitCount         bool     `json:"hitCount"`
	OptimizeCountSql bool     `json:"optimizeCountSql"`
	Orders           []*order `json:"orders"`
}
