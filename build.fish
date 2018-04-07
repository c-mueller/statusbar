#!/usr/bin/env fish
set revision (git rev-parse HEAD)
set rev_count (git rev-list --all --count)
set build_time (date)
set version (git describe --tags --exact-match; or git symbolic-ref -q --short HEAD)
go build -v -ldflags "-X main.version=$version -X main.revision=$revision -X main.buildNumber=$rev_count -X \"main.buildTimestamp=$build_time\""