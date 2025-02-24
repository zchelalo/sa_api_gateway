package classManagementDomain

type ClassEntity struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Grade   string `json:"grade"`
	Code    string `json:"code"`
}
