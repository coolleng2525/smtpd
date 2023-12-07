# go mod edit -replace github.com/mhale/smtpd=../smtpd
script_dir=$(cd $(dirname $0); pwd)

top_dir=$(cd $(dirname $0); cd ../../; pwd)
bin_dir=$top_dir/bin
cert_dir=$top_dir/certs
if [ ! -d $bin_dir ]; then
    mkdir $bin_dir
fi
SMTPD=$bin_dir/smtpd_tls

build() {
    GO111MODULE=auto go build -o $SMTPD main.go 
}
run() {
    $SMTPD -cert $cert_dir/server.pem -key $cert_dir/key.pem -enbaletls
}
cmd=$1
if [ "$cmd" = "" ]; then
    if [ -f $SMTPD ]; then
        run
    else 
        build
        run
    fi
elif [ "$cmd" = "build" ]; then
    build
elif [ "$cmd" = "run" ]; then
   run
elif [ "$cmd" = "test" ]; then
    GO111MODULE=auto go test -v ./...
elif [ "$cmd" = "clean" ]; then
    rm -rf $SMTPD
else
    echo "unknown command: $cmd"
fi
