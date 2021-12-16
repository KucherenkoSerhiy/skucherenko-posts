package service

import "github.com/KucherenkoSerhiy/skucherenko-posts.git/entity"

type CarDetailsService interface {
	GetDetails() entity.CarDetails
}
