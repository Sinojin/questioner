package adapters

import (
	"errors"
	"github.com/sinojin/questioner/internal/questioner/domain/statistic"
	"sync"
)

var ThereIsNoStatistics = errors.New("There is no statistics!")

type statisticRepository struct {
	Statistics []statistic.Statistic
	m          *sync.Mutex
}

func (s *statisticRepository) Get() ([]statistic.Statistic, error) {
	s.m.Lock()
	defer s.m.Unlock()
	if len(s.Statistics) == 0 {
		return nil, ThereIsNoStatistics
	}
	return s.Statistics, nil
}

func (s *statisticRepository) Save(stat statistic.Statistic) error {
	s.m.Lock()
	defer s.m.Unlock()
	s.Statistics = append(s.Statistics, stat)
	return nil
}

func NewStatisticsRepository() statistic.StatisticRepository {
	return &statisticRepository{Statistics: make([]statistic.Statistic, 0), m: &sync.Mutex{}}
}
