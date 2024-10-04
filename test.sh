#!/usr/bin/env bash

cover () {
    local t=$(mktemp -t cover.XXX)
    go test $COVERFLAGS -coverprofile=$t $@
    local testcoverage=$(go tool cover -func=$t | grep total | grep -Eo '[0-9]+\.[0-9]+' | bc)
    unlink $t

    if [ $testcoverage -lt 80 ]; then
        exit 1;
    fi
}

cover