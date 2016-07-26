package dbmgr

// ReturnValue Find等の戻り値
type ReturnValue struct {
	Error error
	Data  []byte
}

// IDB is Interface Data Base
type IDB interface {
	NewIDB()
	Initialize(interface{}) error
	Find(interface{}) *ReturnValue
	Insert(interface{}) *ReturnValue
	Update(interface{}) *ReturnValue
}
