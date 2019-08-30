#build the code for NU client
baseDir=$PWD
export GOPATH=$baseDir
#echo $GOPATH
srcFolder=$baseDir/src
goFiles="$srcFolder/core/*.go $srcFolder/logger/*.go"
#It format the all go files
gofmt -w $goFiles

#It find out style mistakes
golint $srcFolder/core/*.go
if [ $? -ne 0 ];then
echo "go lint errors are there"
exit 1
fi
golint $srcFolder/logger/*.go
if [ $? -ne 0 ];then
echo "build is failed"
exit 1
fi

#Vet examines Go source code and reports suspicious constructs
go vet $srcFolder/core/*.go
if [ $? -ne 0 ];then
echo "build is failed"
exit 1
fi
go vet $srcFolder/logger/*.go
if [ $? -ne 0 ];then
echo "build is failed"
exit 1
fi

#building the https_server code
GOOS=linux GOARCH=amd64 go build -o $baseDir/bin/lottery-system $srcFolder/core/*.go 
if [ $? -ne 0 ];then
echo "build is failed"
exit 1
fi

rm -rf deploy-package
mkdir -p deploy-package
mkdir -p deploy-package/lotterySystem


srcFiles="bin config"
cp -avrdf $srcFiles deploy-package/lotterySystem
