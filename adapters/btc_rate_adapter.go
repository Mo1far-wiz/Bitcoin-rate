package adapters

type APIInterface interface {
	GetBTCRate() (float64, error)
}

type RemoteService struct {
	Remote APIInterface
}

const Currency = "UAH"

func (rs *RemoteService) GetBTCRate() (float64, error) {
	return rs.Remote.GetBTCRate()
}
