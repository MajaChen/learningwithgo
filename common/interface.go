package common

type myinterface interface {
	hello()
}

type yourinterface interface {
	hello()
}

type mi struct {
	myinterface
	yourinterface
}
