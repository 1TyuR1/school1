package service

import (
	"github.com/Fr0olyy/school1/pkg/repository"
)

type Authorization interface {
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
