package conversions

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type L2D struct {
	base.NoOperandsInstruction
}

func (self *L2D) Execution(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	d := float64(l)
	stack.PushDouble(d)
}

type L2F struct {
	base.NoOperandsInstruction
}

func (self *L2F) Execution(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	f := float32(l)
	stack.PushFloat(f)
}

type L2I struct {
	base.NoOperandsInstruction
}

func (self *L2I) Execution(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}
