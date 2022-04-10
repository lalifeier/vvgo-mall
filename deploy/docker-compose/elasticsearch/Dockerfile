ARG ELASTIC_VERSION

FROM docker.elastic.co/elasticsearch/elasticsearch:${ELASTIC_VERSION}
# RUN elasticsearch-plugin install analysis-icu


RUN elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v${ELASTIC_VERSION}/elasticsearch-analysis-ik-6.3.0.zip
