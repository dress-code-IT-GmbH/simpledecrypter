find . -iname "*_test.go" | \
  while read tf ; do
    echo `dirname ${tf}`
  done |
  uniq |
  while read td ; do
    pushd ${td}
    go test -v
    popd
  done