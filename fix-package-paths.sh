#!/bin/bash
PACKAGES=('common' 'enums' 'errors' 'resources' 'services')

function fix_package_path() {
    FILE=$1
    PACKAGE=$2
    MATCH="google.golang.org\/genproto\/googleapis\/ads\/googleads\/v8\/"
    REPLACE="github.com\/hurbcom\/google-ads-go\/"
    sed -i "s/$MATCH$PACKAGE/$REPLACE$PACKAGE/g" $FILE
}

function fix_package_name() {
    FILE=$1
    PACKAGE=$2
    sed -i "s/google_ads_googleads_v8_$PACKAGE/$PACKAGE/g" $FILE
}

echo "fixing packages"
for file in ./google/ads/googleads/v8/**/*.pb.go; do
    for p in "${PACKAGES[@]}"; do
        fix_package_path $file $p
        fix_package_name $file $p
    done
done
mv ./google/ads/googleads/v8/* ./
echo "finished fixing packages"