# these are the values we want to pass for VERSION and BUILD
VERSION=0.1.0
BUILD=`git rev-parse HEAD`

# setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

# default target: builds the project
default:
	go build ${LDFLAGS}

# installs our project: copies binaries
install:
	go install ${LDFLAGS}
	
# cleans our project: deletes binaries
clean:
	go clean

.PHONY: clean install