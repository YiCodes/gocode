package gocode

import "io"

type CodeWriter struct {
	writer       io.Writer
	Indent       int
	unindent     bool
	LineString   string
	IndentString string
}

func NewCodeWriter(writer io.Writer) *CodeWriter {
	w := &CodeWriter{}
	w.writer = writer
	w.unindent = true
	w.LineString = "\r\n"
	w.IndentString = "\t"

	return w
}

func (w *CodeWriter) WriteIndent() {
	for i := 0; i < w.Indent; i++ {
		w.writer.Write([]byte(w.IndentString))
	}
	w.unindent = false
}

func (w *CodeWriter) WriteLine(codes ...string) {
	for _, c := range codes {
		w.Write(c)
	}

	w.writer.Write([]byte(w.LineString))
	w.unindent = true
}

func (w *CodeWriter) Write(code string) {
	if w.unindent {
		w.WriteIndent()
	}

	w.writer.Write([]byte(code))
}

func (w *CodeWriter) BeginBlock(symbol string) {
	if !w.unindent {
		w.writer.Write([]byte(" "))
	}

	w.writer.Write([]byte(symbol))
	w.WriteLine()
	w.Indent++
}

func (w *CodeWriter) EndBlock(symbol string) {
	w.Indent--
	w.Write(symbol)
	w.WriteLine()
}
