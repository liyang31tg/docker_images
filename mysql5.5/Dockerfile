FROM mysql:5.5

RUN mkdir data && mysqldump -ubfbuser -pbfbuserpwd --all-databases > /data/`date +%Y%m%d`.sql && crontab -l 0 4 * * * mysqldump -ubfbuser -pbfbuserpwd --all-databases > /data/`date +%Y%m%d`.sql
