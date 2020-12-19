package format

type DrawableCtr interface {
	Pre()
	Next()
	Get() interface{}
}