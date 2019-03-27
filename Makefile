.PHONY:

NAME = artviz
PWD := $(MKPATH:%/Makefile=%)

### Colour Definitions
END_COLOR=\x1b[0m
GREEN_COLOR=\x1b[32;01m
RED_COLOR=\x1b[31;01m
YELLOW_COLOR=\x1b[33;01m

###
# Go tasks
###
compile-linux:
	@echo "$(GREEN_COLOR)Compiling linux binaries in ./bin $(END_COLOR)"
	CGO_ENABLED=0 GOOS=linux go build -o bin/linux/$(NAME)

compile-win:
	@echo "$(GREEN_COLOR)Compiling win binaries in ./bin $(END_COLOR)"
	CGO_ENABLED=0 GOOS=windows go build -o bin/win/$(NAME).exe

compile-mac:
	@echo "$(GREEN_COLOR)Compiling mac binaries in ./bin $(END_COLOR)"
	CGO_ENABLED=0 GOOS=darwin go build -o bin/mac/$(NAME)

compile: compile-mac compile-win compile-linux

###
# Docker tasks
###
compose-up:
	@POSTGRES_PSWRD="password" docker-compose -f artifactory-pro.yml up -d

compose-stop:
	@docker-compose -f artifactory-pro.yml stop

compose-restart:
	@docker-compose -f artifactory-pro.yml restart

compose-ps:
	@docker-compose -f artifactory-pro.yml ps

compose-logs:
	@docker-compose -f artifactory-pro.yml logs

compose-rm:
	@docker-compose -f artifactory-pro.yml rm