package dto


type Status struct{
	Code int `json:"code"`
	Message string `json:"message"`
}
type SingleResponse struct{
	Status Status `json:"status"`
	Data any `json:"data"`
}
type Paging struct{
	Page int `json:"page"`
	Size int `json:"size"`
	TotalRows int `json:"totalRows"`
	TotalPage int `json:"totalPages"`
}
type PagingResponse struct{
	Status Status `json:"status"`
	Data []any `json:"data"`
	Paging Paging `json:"paging"`
}

// {
// 	"status": {
// 		"code":200,
// 		"message":"success create data"
// 	},
// 	data :{

// 	}
// }