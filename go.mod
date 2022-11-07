module sample_todo_api

go 1.19

require (
	github.com/go-sql-driver/mysql v1.6.0
	local.packages/model/entity v0.0.0
	local.packages/model/repository v0.0.0
	local.packages/controller/dto v0.0.0
// local.packages/controller v0.0.0
)

replace local.packages/model/entity => ./model/entity

replace local.packages/model/repository => ./model/repository

replace local.packages/controller/dto => ./controller/dto

// replace local.packages/controller => ./controller
