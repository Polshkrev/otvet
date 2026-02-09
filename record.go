package otvet

// Standardized API response record.
type Record[Type any] struct {
	Code StatusCode `json:"code"`
	Data Type       `json:"data"`
}

// Construct a new record.
// Returns a new record with a given status code and a given payload of data.
func NewRecord[Type any](code StatusCode, data Type) *Record[Type] {
	var record *Record[Type] = new(Record[Type])
	record.Code = code
	record.Data = data
	return record
}
