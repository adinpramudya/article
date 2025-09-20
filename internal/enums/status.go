package enums

const (
	StatusPublish = "Publish"
	StatusDraft   = "Draft"
	StatusThrash  = "Thrash"
)


// Daftar valid status
var ValidStatuses = []string{
	StatusPublish,
	StatusDraft,
	StatusThrash,
}

// Fungsi helper untuk validasi status
func IsValidStatus(status string) bool {
	for _, v := range ValidStatuses {
		if v == status {
			return true
		}
	}
	return false
}