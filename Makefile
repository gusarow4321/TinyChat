VERSION=$(shell git describe --tags --always)
DANGLING=$(shell docker images --filter "dangling=true" -q --no-trunc)

auth-img:
	docker image build --build-arg VERSION=$(VERSION) ./auth -t tiny-chat-auth:$(VERSION)

gw-img:
	docker image build --build-arg VERSION=$(VERSION) ./gateway -t tiny-chat-gateway:$(VERSION)

msg-img:
	docker image build --build-arg VERSION=$(VERSION) ./messenger -t tiny-chat-messenger:$(VERSION)

all-imgs:
	make auth-img
	make gw-img
	make msg-img

rm-dangling:
	docker rmi $(DANGLING)

auth-test:
	cd auth && go test -coverprofile=auth-coverage.txt -v ./internal/...

messenger-test:
	cd messenger && go test -coverprofile=messenger-coverage.txt -v ./internal/...

test:
	make auth-test
	make messenger-test
