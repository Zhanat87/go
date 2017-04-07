FROM phpdockerio/php71-fpm:latest
#  php7.1-mongo
# Install selected extensions and other stuff
RUN apt-get update; \
    apt-get -y --no-install-recommends install git \
    php-memcached php7.1-mysql php7.1-pgsql php7.1-mbstring php7.1-xml php7.1-curl; \
    apt-get clean; rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* /usr/share/doc/* ; \
    curl -sL https://deb.nodesource.com/setup_7.x | bash - ; \
    apt-get update; apt-get -y --no-install-recommends install nodejs build-essential;
#    /usr/bin/npm install gulp -g

WORKDIR "/var/www/stack-frontend"
