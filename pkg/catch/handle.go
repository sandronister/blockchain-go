package catch

func Handle(err error) {
	if err != nil {
		panic(err)
	}
}
