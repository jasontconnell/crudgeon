@echo OFF

go build ./cmd/crudgeon

crudgeon -dir ./datafiles -path output -obj Business -ns Name.Space
crudgeon -dir ./datafiles -path output -obj Employee -ns Name.Space