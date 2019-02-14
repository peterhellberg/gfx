// +build js

package gfx

import "syscall/js"

// JS gives access to js.Global and js.TypedArrayOf
var JS = JavaScript{
	Global:       js.Global,
	TypedArrayOf: js.TypedArrayOf,
}

// JavaScript is a type that contains fields with Global and TypedArrayOf funcs.
type JavaScript struct {
	Global       func() js.Value
	TypedArrayOf func(interface{}) js.TypedArray
}
