#!/bin/bash

gofiles=$(find . -name "*.go" | grep -v "/gen/")

for gofile in $gofiles; do
	echo $gofile
	sed '/^import/,/^[[:space:]]*)/ { /^[[:space:]]*$/ d; }' $gofile >tmp
	mv tmp $gofile

	goimports -local github.com/hiromaily/ -w $gofile
done

echo "Done imports"
