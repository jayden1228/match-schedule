package model

//BaseModel BaseModel
type BaseModel struct {
	ID        int64 `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	CreatedAt int64 `gorm:"column:created_at"`
	UpdatedAt int64 `gorm:"column:updated_at"`
}

//APIResponse APIResponse
type APIResponse struct {
	Code    int         `json:"code" binding:"required"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

//PageResponseModel PageResponseModel
type PageResponseModel struct {
	List       interface{} `json:"list"`
	PageNumber int         `json:"pageNumber"`
	PageSize   int         `json:"pageSize"`
	TotalCount int         `json:"totalCount"`
	TotalPage  int         `json:"totalPage"`
}

//PageRequestModel PageRequestModel
type PageRequestModel struct {
	PageNumber int    `json:"pageNumber"`
	PageSize   int    `json:"pageSize"`
	Sort       string `json:"sort"`
	Order      string `json:"order"`
}
