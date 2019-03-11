package freerangeserver

import (
	"testing"

	"github.com/robertkrimen/otto"
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

func TestCallGoFromLua(t *testing.T) {

	callCount := 0
	Add := func(L *lua.LState) int {
		a := L.Get(1).(lua.LNumber)
		b := L.Get(2).(lua.LNumber)
		L.Push(a + b)
		callCount++
		return 1
	}

	L := lua.NewState()
	defer L.Close()
	L.SetGlobal("CallAdd", L.NewFunction(Add)) /* Original lua_setglobal uses stack... */
	if err := L.DoString(`CallAdd(1,2)`); err != nil {
		panic(err)
	}
	if callCount != 1 {
		t.Error("expected a call")
	}
}

func TestOttoBasic(t *testing.T) {
	vm := otto.New()
	vm.Run(`
		abc = 2 + 2;
		console.log("The value of abc is " + abc); // 4
	`)
	if value, err := vm.Get("abc"); err == nil {
		if value_int, err := value.ToInteger(); err == nil {
			if value_int != 4 {
				t.Error("not 4")
			}
		}
	}
}

func TestGoCallFromOtto(t *testing.T) {
	vm := otto.New()
	vm.Set("twoPlus", func(call otto.FunctionCall) otto.Value {
		right, _ := call.Argument(0).ToInteger()
		result, _ := vm.ToValue(2 + right)
		return result
	})

	result, _ := vm.Run(`
    
    result = twoPlus(2.0); // 4
	`)
	i, _ := result.ToInteger()
	if i != 4 {
		t.Error("not 4")
	}

}