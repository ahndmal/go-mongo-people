startdb:
	@harness/setup.sh start

stopdb:
	@harness/setup.sh stop
deps:
	go get github.com/qiniu/qmgo
	go get github.com/julienschmidt/httprouter
	go get go.mongodb.org/mongo-driver