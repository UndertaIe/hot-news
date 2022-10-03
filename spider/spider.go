package spider

type Spider struct {
	Name string
	middlewares []middleware
}

func (s *Spider) Use(middleware ...string) {

}
