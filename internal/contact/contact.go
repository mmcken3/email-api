package contact

// Contact will represent a contact structure.
type Contact struct {
	Name    string `json:"name"`
	Email   string `json:"email_address"`
	Message string `json:"message"`
}
