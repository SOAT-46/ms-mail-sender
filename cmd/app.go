package main

import "github.com/soat-46/ms-mail-sender/internal/global/domain/entities"

func RunModules(apps []entities.App) {
	for _, module := range apps {
		module.RunConsumers()
	}
}
