package data

type CsType int
type SqlType int

const (
	CNone CsType = iota
	CInt
	CDecimal
	CDouble
	CString
	CShort
	CLong
	CDateTime
)

const (
	SNone SqlType = iota
	SInt
	SDecimal
	SDouble
	SString
	SShort
	SLong
	SDateTime
)

type Field struct {
	Name        string
	Type        string
	CsNullable  bool
	SqlNullable bool
	CsType      CsType
	SqlType     SqlType
}
