gen-boiler:
	cd sqlboiler && \
	sqlboiler mysql -c sqlboiler.toml -o entity -p sqlboiler --no-tests --wipe --templates /$(GOPATH)/pkg/mod/github.com/volatiletech/sqlboiler/v4@v4.16.2/templates/main /$(GOPATH)/pkg/mod/github.com/volatiletech/sqlboiler/v4@v4.16.2/templates/test --templates ./templates

gen-mock:
	@for file in $(shell ls dao/*.go); do \
		mockgen -source=$$file -destination=dao/mock$${file/dao/} -package=mockdao && \
		echo "Generated mock for $$file"; \
	done
	@mockgen -source=usecase/score.go -destination=usecase/mock/score.go -package=mockusecase && \
	echo "Generated mock for usecase/score.go"

gen-open-api:
	cd controller && \
	oapi-codegen -generate types,gin,spec,skip-prune -package controller -o ./schema.gen.go ./../../open-api/open-api.yaml

gen-all: gen-boiler gen-mock gen-open-api
