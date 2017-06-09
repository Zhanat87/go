FROM docker.elastic.co/kibana/kibana:5.2.1

# add our user and group first to make sure their IDs get assigned consistently
# RUN groupadd -r kibana && useradd -r -m -g kibana kibana

#ENV KIBANA_HOME /opt/kibana
# RUN mkdir ${KIBANA_HOME} && mkdir -p /var/log/kibana && \
#     chown -R kibana:kibana ${KIBANA_HOME} /var/log/kibana;
#WORKDIR ${KIBANA_HOME}
#USER kibana
#RUN bin/kibana plugin --install elastic/sense
#RUN bin/kibana plugin --install elasticsearch/marvel/latest

# ENV ELASTICSEARCH_URL http://172.17.0.3:9200