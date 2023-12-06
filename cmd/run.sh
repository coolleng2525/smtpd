# go mod edit -replace github.com/mhale/smtpd=../smtpd
script_dir=$(cd $(dirname $0); pwd)

top_dir=$(cd $(dirname $0); cd ..; pwd)
bin_dir=$top_dir/bin
if [ ! -d $bin_dir ]; then
    mkdir $bin_dir
fi
SMTPD=$bin_dir/smtpd

build() {
    GO111MODULE=auto go build -o $SMTPD server.go 
}
cmd=$1
if [ "$cmd" = "" ]; then
    if [ -f $SMTPD ]; then
        $SMTPD
    else 
        build
        $SMTPD
    fi
elif [ "$cmd" = "build" ]; then
    build
elif [ "$cmd" = "run" ]; then
    $SMTPD
elif [ "$cmd" = "test" ]; then
    GO111MODULE=auto go test -v ./...
elif [ "$cmd" = "clean" ]; then
    rm -rf $SMTPD
else
    echo "unknown command: $cmd"
fi
