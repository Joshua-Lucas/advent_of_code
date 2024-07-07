#Makefile

# Declare variables for flags
YEAR = 2015
DAY = 1

# Run to get go solutions. 
.PHONY: run-go
run-go: 
	cd go && go run . -year=$(YEAR) -day=$(DAY)


# Run to get go solutions. 
.PHONY: run-rust
run-rust: 
	cd rust && cargo run -- --year $(YEAR) --day $(DAY)  

