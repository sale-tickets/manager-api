mkdir -p docs
rm -rf docs/*
protoc \
    -I=. \
    --openapiv2_out=docs \
    --openapiv2_opt logtostderr=true \
    api/*.proto