package database

// Code generated by chromedp-gen. DO NOT EDIT.

// ID unique identifier of Database object.
type ID string

// String returns the ID as string value.
func (t ID) String() string {
	return string(t)
}

// Database database object.
type Database struct {
	ID      ID     `json:"id"`      // Database ID.
	Domain  string `json:"domain"`  // Database domain.
	Name    string `json:"name"`    // Database name.
	Version string `json:"version"` // Database version.
}

// Error database error.
type Error struct {
	Message string `json:"message"` // Error message.
	Code    int64  `json:"code"`    // Error code.
}
