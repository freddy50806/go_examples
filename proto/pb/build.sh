BASEDIR=$(dirname "$0")
echo $BASEDIR;
protoc -I $BASEDIR $BASEDIR/*.proto --go_out=plugins=grpc:$BASEDIR/golang