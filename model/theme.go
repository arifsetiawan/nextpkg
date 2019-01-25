package model

// CustomTheme ...
type CustomTheme struct {
	ID       string   `json:"id"`
	TenantID string   `json:"tenant_id"`
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	Title    string   `json:"title"`
	LogoURL  string   `json:"logo_url"`
	StyleURL string   `json:"style_url"`
	FileURLs []string `json:"file_urls"`
}
