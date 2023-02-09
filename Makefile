## test: test all the files in the current directory and detector race | coverage
test: Makefile
	go test ./... -v -race -cover

help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'