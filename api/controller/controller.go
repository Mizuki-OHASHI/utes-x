//go:generate oapi-codegen -generate types,gin,spec,skip-prune -package controller -o ./schema.gen.go ./../../open-api/open-api.yaml
package controller
