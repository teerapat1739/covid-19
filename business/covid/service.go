package covid

type Service interface {
	CovidSummary() (interface{}, error)
}
