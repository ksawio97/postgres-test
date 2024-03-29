FROM postgres:16.2-alpine

ENV POSTGRES_DB=test
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres

COPY create_table.sql /docker-entrypoint-initdb.d/

EXPOSE 5432