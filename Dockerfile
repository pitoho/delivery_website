FROM postgres:14.5-alpine

COPY db.sql /docker-entrypoint-initdb.d/

CMD ["postgres"]