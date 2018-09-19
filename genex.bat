@echo OFF

go build

crudgeon -file datafiles\example1.txt -path output -obj Business
crudgeon -file datafiles\example2.txt -path output -obj Employee