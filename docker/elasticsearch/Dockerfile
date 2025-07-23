FROM docker.elastic.co/elasticsearch/elasticsearch:9.0.4

RUN bin/elasticsearch-plugin install analysis-kuromoji \
  && bin/elasticsearch-plugin install analysis-nori
