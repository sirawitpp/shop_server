package service

import "sirawit/shop/internal/model"

type LoggerService interface {
	WriteLogLoginTimestamp(input model.Logger) (*model.Logger, error)
}

type loggerService struct{}

func NewLoggerService() LoggerService {
	return &loggerService{}
}

func (l *loggerService) WriteLogLoginTimestamp(input model.Logger) (*model.Logger, error) {
	return &input, nil
}
