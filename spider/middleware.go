package spider

func init() {

}

type middleware interface {
	before(r *Request) *Request
	After()
}

func registerMiddlewares() {

}

func PrintMiddlewares() {

}
