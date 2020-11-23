all: install test

install:
	go install .

test: clean
	genny new github.com/petomalina/genny/test --service-dir=src && \
	cd test && \
	genny add proto github.com/googleapis/api-common-protos && \
	genny add api health/v1 && \
	genny add api monitoring/dashboard/v1

clean:
	rm -rf test/