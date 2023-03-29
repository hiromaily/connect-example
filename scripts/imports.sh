#!/bin/bash

gofiles=$(find . -name "*.go" | grep -v "/gen/")

for gofile in $gofiles; do
	echo $gofile
	sed '/^import/,/^[[:space:]]*)/ { /^[[:space:]]*$/ d; }' $gofile >tmp
	mv tmp $gofile
done

go fmt $(go list ./... | grep -v "/gen/")
goimports -local github.com/hiromaily/ -w $(goimports -local github.com/hiromaily/ -l ./ | grep -v "/vendor/")
