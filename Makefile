help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

startSQL: ## Starts the PostgreSQL docker image. Will fail. 2nd command must be run in bash.
	sudo docker run -p 8081:5432 --rm --name accountsql -v "$(PWD)/internal/database/:/opt/accountsql/" -e POSTGRES_PASSWORD=pass -e POSTGRES_DB=GoKitchen -d postgres
	sudo docker exec -it accountsql psql -U postgres -f /opt/accountsql/startup.sql
	##sudo docker exec -it accountsql psql -U postgres -c ""

stopSQL: ## Stops the sql container
	sudo docker stop accountsql
