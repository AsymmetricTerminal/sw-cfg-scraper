#!/usr/bin/env bash
package=$1
if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi

platforms=("windows/amd64" "linux/amd64" "darwin/amd64")
for platform in "${platforms[@]}"
do
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name=$package_name'-'$GOOS'-'$GOARCH

	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi
	env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name $package
done
