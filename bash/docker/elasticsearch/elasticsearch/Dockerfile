FROM elasticsearch:2.4

RUN bin/plugin install mobz/elasticsearch-head
RUN bin/plugin install license
RUN bin/plugin install marvel-agent

# Mount elasticsearch.yml config
# move to /usr/share/elasticsearch/config/elasticsearch.yml
# ADD elasticsearch/config/elasticsearch.yml config/elasticsearch.yml

WORKDIR "/usr/share/elasticsearch"
