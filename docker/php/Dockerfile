# https://hub.docker.com/_/php/
FROM ubuntu:latest

RUN apt-get update && apt-get install -y php php-pgsql

ADD test.php test.php

CMD [ "php", "test.php" ]
