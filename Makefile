DANGLING=$(shell docker images --filter "dangling=true" -q --no-trunc)

authimg:
	docker image build ./auth -t tiny-chat-auth

gwimg:
	docker image build ./gateway -t tiny-chat-gateway

msgimg:
	docker image build ./messenger -t tiny-chat-messenger

allimgs:
	make authimg
	make gwimg
	make msgimg

rmdangling:
	docker rmi $(DANGLING)