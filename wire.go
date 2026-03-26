//go:build wireinject
// +build wireinject

package main

import (
	"ems/handler"
	"ems/repository"
	"ems/services"

	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitializeEmployeeHandler(pool *pgxpool.Pool) *handler.EmployeeHandler {
	wire.Build(
		repository.NewEmployeeRepository,
		services.NewEmployeeService,
		handler.NewEmployeeHandler,
	)
	return &handler.EmployeeHandler{}
}

func InitializeDepartmentHandler(pool *pgxpool.Pool) *handler.DepartmentHandler {
	wire.Build(
		repository.NewDepartment,
		services.NewDepartmentService,
		handler.NewDepartmentHandler,
	)
	return &handler.DepartmentHandler{}
}

func InitializeAssetHandler(pool *pgxpool.Pool) *handler.AssetHandler {
	wire.Build(
		repository.NewAssetRepository,
		services.NewAssetService,
		handler.NewAssetHandler,
	)
	return &handler.AssetHandler{}
}

func InitializeSalaryHandler(pool *pgxpool.Pool) *handler.SalaryCategoryHandler {
	wire.Build(
		repository.NewSalaryCategoryRepository,
		services.NewSalaryCategoryService,
		handler.NewSalaryCategoryHandler,
	)
	return &handler.SalaryCategoryHandler{}
}