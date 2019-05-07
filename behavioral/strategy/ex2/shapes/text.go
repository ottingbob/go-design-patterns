package shapes

import "example2"

type TextSqaure struct {
	strategy.DrawOutput
}

func (t *TextSquare) Draw() error {
	t.Writer.Write([]byte("Circle"))
	return nil
}