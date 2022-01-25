FROM mysql:5.7

COPY ./ddl /docker-entrypoint-initdb.d/
COPY ./conf/my.cnf /etc/mysql/my.cnf
ENV MYSQL_ROOT_PASSWORD=codehanhan