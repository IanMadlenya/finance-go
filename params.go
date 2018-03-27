package finance

import "context"

const (
	endbefore  = "ending_before"
	startafter = "starting_after"
)

// Params is the structure that contains the common properties
// of any *Params structure.
type Params struct {
	// Context used for request. It may carry deadlines, cancelation signals,
	// and other request-scoped values across API boundaries and between
	// processes.
	//
	// Note that a cancelled or timed out context does not provide any
	// guarantee whether the operation was or was not completed.
	Context context.Context `form:"-"`

	End   string   `form:"ending_before"`
	Exp   []string `form:"expand"`
	Limit int      `form:"limit"`

	// Single specifies whether this is a single page iterator. By default,
	// listing through an iterator will automatically grab additional pages as
	// the query progresses. To change this behavior and just load a single
	// page, set this to true.
	Single bool `form:"-"` // Not an API parameter

	Start string `form:"starting_after"`
}

// ListMeta is the structure that contains the common properties
// of List iterators. The Count property is only populated if the
// total_count include option is passed in (see tests for example).
type ListMeta struct {
	Count uint32 `json:"total_count"`
	More  bool   `json:"has_more"`
	URL   string `json:"url"`
}
