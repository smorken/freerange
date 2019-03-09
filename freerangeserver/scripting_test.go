package freerangeserver

import "testing"

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/yuin/gopher-lua"
)

func TestLuaBasic(t *testing.T) {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile("add.lua"); err != nil {
		panic(err)
	}
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("add2"),
		NRet:    1,
		Protect: true,
	}, lua.LNumber(10), lua.LNumber(20)); err != nil {
		panic(err)
	}
	ret := L.Get(-1).(lua.LNumber) // returned value
	if ret != 30 {
		t.Error("incorrect number")
	}
	L.Pop(1) // remove received value
}

func TestGopherjsBasic(t *testing.T) {

	js.Global.Set("a", 1)
	result := js.Global.Get("a").Float()
	if result != 1.0 {
		t.Error("unexpected result")
	}
}

// func TestGojaBasic(t *testing.T) {
// 	vm := goja.New()
// 	v, err := vm.RunString("2 + 2")
// 	if err != nil {
// 		panic(err)
// 	}
// 	if num := v.Export().(int64); num != 4 {
// 		panic(num)
// 	}
// }
