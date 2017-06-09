FROM elasticsearch:latest

#RUN bin/plugin install mobz/elasticsearch-head

# note: before replace original config file, need copy original config file and custom config it
# Mount elasticsearch.yml config
# move to /usr/share/elasticsearch/config/elasticsearch.yml
COPY ./config/elasticsearch.yml /usr/share/elasticsearch/config/elasticsearch.yml

WORKDIR "/usr/share/elasticsearch"
