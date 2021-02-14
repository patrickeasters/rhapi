generate:
	oapi-codegen -config rhsm/codegen-config.yml specs/rhsm-v1.yml
	oapi-codegen -config sources/codegen-config.yml specs/sources-v3.yml

