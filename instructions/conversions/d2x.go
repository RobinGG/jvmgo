package conversions

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type D2F struct {
	base.NoOperandsInstruction
}

func (self *D2F) Execution(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

type D2I struct {
	base.NoOperandsInstruction
}

func (self *D2I) Execution(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

type D2L struct {
	base.NoOperandsInstruction
}

func (self *D2L) Execution(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}
