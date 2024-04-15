#Makefile

# Declare variables for flags
YEAR = 2015
DAY = 1

# Run to get go solutions. 
.PHONY: go
go: 
	cd go && go run . -year=$(YEAR) -day=$(DAY)
