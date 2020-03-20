package octopusdeploy

type PagedResults struct {
	ItemType       string `json:"ItemType"`
	TotalResults   int    `json:"TotalResults"`
	NumberOfPages  int    `json:"NumberOfPages"`
	LastPageNumber int    `json:"LastPageNumber"`
	ItemsPerPage   int    `json:"ItemsPerPage"`
	IsStale        bool   `json:"IsStale"`
	Links          Links  `json:"Links"`
}

type Links struct {
	Self        string `json:"Self"`
	Template    string `json:"Template"`
	PageAll     string `json:"Page.All"`
	PageCurrent string `json:"Page.Current"`
	PageLast    string `json:"Page.Last"`
	PageNext    string `json:"Page.Next"`
}

type User struct {
	ID                  string `json:"Id"`
	Username            string `json:"Username"`
	DisplayName         string `json:"DisplayName"`
	IsActive            bool   `json:"IsActive"`
	IsService           bool   `json:"IsService"`
	EmailAddress        string `json:"EmailAddress"`
	CanPasswordBeEdited bool   `json:"CanPasswordBeEdited"`
	IsRequestor         bool   `json:"IsRequestor"`
	Links               struct {
		Self        string `json:"Self"`
		Permissions string `json:"Permissions"`
		APIKeys     string `json:"ApiKeys"`
		Avatar      string `json:"Avatar"`
	} `json:"Links"`
}
