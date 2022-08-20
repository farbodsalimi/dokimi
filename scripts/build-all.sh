#!/bin/bash

# check arguments
package=$1
if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi

# clean up bin directory
rm -rf bin

# parse aerguments
package_split=(${package//\// })
package_name=${package_split[2]}

# spesify targeted platforms
platforms=("linux/amd64" "darwin/amd64" "windows/amd64")
version=$(git sv nv)

for platform in "${platforms[@]}"
do
  platform_split=(${platform//\// })
  platform_name=${platform_split[0]}
  platform_arch=${platform_split[1]}

  output_path=bin/$package_name'_'$version'-'$platform_name'-'$platform_arch
  if [ $platform_name = "windows" ]; then
    output_path+='.exe'
  fi

  # compile
  env GOOS=$platform_name GOARCH=$platform_arch go build -ldflags="-X 'main.Version=v${version}'" -o $output_path $package

  # compress binaries
  if [ $platform_name = "windows" ]; then
    zip -j $output_path.zip $output_path
    rm $output_path
  else
    tar -czf $output_path.tar.gz $output_path
    rm $output_path
  fi

  if [ $? -ne 0 ]; then
    echo 'An error has occurred! Aborting the script execution...'
    exit 1
  fi
done
