URL="https://www.chiikinogennki.soumu.go.jp/k-cloud-api/search/publicFiles/kanko_all.csv"
curl -s --head $URL | grep Last-Modified > build/Last-Modified
diff build/Last-Modified docs/Last-Modified
ret=$?
if [ $ret -eq 0 ]; then
    echo "don't need to do"
    exit 0
fi
curl -o build/kanko_all.csv $URL
cp -f build/Last-Modified docs/Last-Modified
./csv2json build/kanko_all.csv > docs/kanko_all.json
git add -A docs/*
git commit -m $(date +%Y%m%d%H%M%S)