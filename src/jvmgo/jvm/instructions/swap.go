package instructions

import "jvmgo/jvm/rtda"

// Swap the top two operand stack values
type swap struct{ NoOperandsInstruction }

func (self *swap) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.Pop()
	val2 := stack.Pop()
	stack.Push(val1)
	stack.Push(val2)
}
