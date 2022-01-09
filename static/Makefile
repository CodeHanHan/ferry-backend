up:
	@docker run --name ferry-static --restart=always \
		-p 10081:80 \
	   	-v ~/.ferry-data/nginx/data:/usr/share/nginx \
		-v ~/.ferry-data/nginx/log:/var/log/nginx \
		-d ferry-nginx

build:
	@docker build -t ferry-nginx .

config-up:
	@docker cp ./nginx-conf/nginx.conf ferry-static:/etc/nginx/conf.d/default.conf
	# @docker restart ferry-static
