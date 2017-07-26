package profiler

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	cdp "github.com/knq/chromedp/cdp"
	"github.com/knq/chromedp/cdp/debugger"
)

// EventConsoleProfileStarted sent when new profile recording is started
// using console.profile() call.
type EventConsoleProfileStarted struct {
	ID       string             `json:"id"`
	Location *debugger.Location `json:"location"`        // Location of console.profile().
	Title    string             `json:"title,omitempty"` // Profile title passed as an argument to console.profile().
}

// EventConsoleProfileFinished [no description].
type EventConsoleProfileFinished struct {
	ID       string             `json:"id"`
	Location *debugger.Location `json:"location"` // Location of console.profileEnd().
	Profile  *Profile           `json:"profile"`
	Title    string             `json:"title,omitempty"` // Profile title passed as an argument to console.profile().
}

// EventTypes all event types in the domain.
var EventTypes = []cdp.MethodType{
	cdp.EventProfilerConsoleProfileStarted,
	cdp.EventProfilerConsoleProfileFinished,
}