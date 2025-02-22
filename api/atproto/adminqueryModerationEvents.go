// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.admin.queryModerationEvents

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// AdminQueryModerationEvents_Output is the output of a com.atproto.admin.queryModerationEvents call.
type AdminQueryModerationEvents_Output struct {
	Cursor *string                   `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
	Events []*AdminDefs_ModEventView `json:"events" cborgen:"events"`
}

// AdminQueryModerationEvents calls the XRPC method "com.atproto.admin.queryModerationEvents".
//
// includeAllUserRecords: If true, events on all record types (posts, lists, profile etc.) owned by the did are returned
// sortDirection: Sort direction for the events. Defaults to descending order of created at timestamp.
// types: The types of events (fully qualified string in the format of com.atproto.admin#modEvent<name>) to filter by. If not specified, all events are returned.
func AdminQueryModerationEvents(ctx context.Context, c *xrpc.Client, createdBy string, cursor string, includeAllUserRecords bool, limit int64, sortDirection string, subject string, types []string) (*AdminQueryModerationEvents_Output, error) {
	var out AdminQueryModerationEvents_Output

	params := map[string]interface{}{
		"createdBy":             createdBy,
		"cursor":                cursor,
		"includeAllUserRecords": includeAllUserRecords,
		"limit":                 limit,
		"sortDirection":         sortDirection,
		"subject":               subject,
		"types":                 types,
	}
	if err := c.Do(ctx, xrpc.Query, "", "com.atproto.admin.queryModerationEvents", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
