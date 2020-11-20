package errcode

type Code int

const (
	Success Code = 0
	ErrNone Code = Success

	//Test error
	ErrInvalidParam Code = 10000

	//authorized error
	ErrAuthInvalidSession Code = 11000
	ErrAuthHeaderNoToken  Code = 11001
	ErrAuthInvalidToken   Code = 11002
	ErrAuthPwdNotMatch    Code = 11004
	ErrAuthUserIsNotExist Code = 11005

	//file error
	ErrFileNotFound Code = 12000

	//project error
	ErrProjectDataBase Code = 13001

	//job error
	ErrJobScriptError Code = 14000
	ErrJobDataBase    Code = 14001
	ErrJobSchedule    Code = 14002
	ErrJobPltPath     Code = 14003

	//gas error
	ErrGasDataBase Code = 15000
)
