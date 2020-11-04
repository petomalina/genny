all: install test

install:
	go install .

test:
	genny new test && \
	cd test && \
	genny add proto github.com/googleapis/api-common-protos && \
	genny add api health/v1 && \
	genny add api monitoring/dashboard/v1