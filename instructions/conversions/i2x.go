package conversions

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type I2D struct {
	base.NoOperandsInstruction
}

func (self *I2D) Execution(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	d := float64(i)
	stack.PushDouble(d)
}

type I2F struct {
	base.NoOperandsInstruction
}

func (self *I2F) Execution(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	f := float32(i)
	stack.PushFloat(f)
}

type I2L struct {
	base.NoOperandsInstruction
}

func (self *I2L) Execution(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	l := int64(i)
	stack.PushLong(l)
}
