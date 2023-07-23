package clockogo

type Paging struct {
	ItemsPerPage int `json:"items_per_page"`
	CurrentPage  int `json:"current_page"`
	CountPages   int `json:"count_pages"`
	CountItems   int `json:"count_items"`
}
