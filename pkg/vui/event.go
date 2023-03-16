package vui

type EventTarger interface {
	addEventListener()
	dispatchEvent()
	removeEventListener()
}
