package statistic

type StatisticRepository interface {
	Get() ([]Statistic, error)
	Save(Statistic) error
}
