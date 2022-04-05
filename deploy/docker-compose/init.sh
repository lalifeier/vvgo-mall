#!/bin/sh


docker exec -it mysql mysql -uroot -p
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' ;
FLUSH PRIVILEGES;

# CREATE USER canal IDENTIFIED BY  'canal';
# GRANT SELECT, REPLICATION SLAVE,  REPLICATION CLIENT ON *.* TO 'canal'@'%';
# # -- GRANT ALL PRIVILEGES ON *.* TO 'canal'@'%' ;
# FLUSH PRIVILEGES;


# kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 -partitions 1 --topic {topic}

sysctl -w vm.max_map_count=262144

# groupadd elsearch
# useradd elsearch -g elsearch
# chown -R elsearch:elsearch data/elk/elasticsearch

# /usr/share/elasticsearch/bin/elasticsearch-setup-passwords interactive
# /usr/share/elasticsearch/bin/elasticsearch-setup-passwords auto
# elastic、apm_system、kibana_system、logstash_system、beats_system、remote_monitoring_user

# docker-compose exec elasticsearch bin/elasticsearch-reset-password --batch --user elastic
docker-compose exec elasticsearch bin/elasticsearch-reset-password --batch --user kibana_system
docker-compose exec elasticsearch bin/elasticsearch-reset-password --batch --user logstash_system

# docker exec elasticsearch /bin/bash -c "bin/elasticsearch-setup-passwords auto --batch --url http://elasticsearch:9200"
# Changed password for user apm_system
# PASSWORD apm_system = 2UWUpNJZjdN5lGOzjZxu

# Changed password for user kibana_system
# PASSWORD kibana_system = kYjkSpvEksACuFhSKAp2

# Changed password for user kibana
# PASSWORD kibana = kYjkSpvEksACuFhSKAp2

# Changed password for user logstash_system
# PASSWORD logstash_system = bUMSmD6uRBBWPDTcpwHh

# Changed password for user beats_system
# PASSWORD beats_system = vpCTUG6zx8ZSEkNIVxdx

# Changed password for user remote_monitoring_user
# PASSWORD remote_monitoring_user = 1E4MXgxsbQZJPLXwhAne

# Changed password for user elastic
# PASSWORD elastic = KXgpPsOIHXJ1ruQOyIkO
