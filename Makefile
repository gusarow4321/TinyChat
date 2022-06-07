VERSION=$(shell git describe --tags --always)
DANGLING=$(shell docker images --filter "dangling=true" -q --no-trunc)

authimg:
	docker image build --build-arg VERSION=$(VERSION) ./auth -t tiny-chat-auth:$(VERSION)

gwimg:
	docker image build --build-arg VERSION=$(VERSION) ./gateway -t tiny-chat-gateway:$(VERSION)

msgimg:
	docker image build --build-arg VERSION=$(VERSION) ./messenger -t tiny-chat-messenger:$(VERSION)

allimgs:
	make authimg
	make gwimg
	make msgimg

rmdangling:
	docker rmi $(DANGLING)