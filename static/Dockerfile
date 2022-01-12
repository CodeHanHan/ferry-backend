FROM nginx:1.21

COPY ./nginx-conf/nginx.conf /etc/nginx/conf.d/default.conf

RUN chmod -R 666 /usr/share/nginx/

EXPOSE 80
