package data

type CodeType int
type SqlType int

const (
	CNone CodeType = iota
	CInt
	CDecimal
	CDouble
	CString
	CShort
)

const (
	SNone SqlType = iota
	SInt
	SDecimal
	SDouble
	SString
	SShort
)

type Field struct {
	Name     string
	Type     string
	Nullable bool
	CodeType CodeType
	SqlType  SqlType
}
