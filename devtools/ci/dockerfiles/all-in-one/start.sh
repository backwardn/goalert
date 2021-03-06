#!/bin/sh
set -e

export PGDATA=/var/lib/postgresql/data PGUSER=postgres DB_URL=postgresql://postgres@

if ! test -f /init-db
then
  mkdir -p ${PGDATA} /run/postgresql /var/log/postgresql &&\
  chown postgres ${PGDATA} /run/postgresql /var/log/postgresql &&\
  su postgres -s /bin/sh -c "initdb $PGDATA" &&\
  echo "host all  all    0.0.0.0/0  md5" >> $PGDATA/pg_hba.conf &&\
  echo "listen_addresses='*'" >> $PGDATA/postgresql.conf &&\
  echo "fsync = off" >> $PGDATA/postgresql.conf &&\
  echo "full_page_writes = off" >> $PGDATA/postgresql.conf
  touch /init-db
fi

su postgres -s /bin/sh -c "pg_ctl start -w -l /var/log/postgresql/server.log"

if ! test -f /init-data
then
  echo Seeding DB with demo data...
  /bin/resetdb -with-rand-data -admin-id=00000000-0000-0000-0000-000000000000 -db-url "$DB_URL" -skip-drop
  /bin/goalert add-user --user-id=00000000-0000-0000-0000-000000000000 --user admin --pass admin123 --db-url "$DB_URL"
  touch /init-data
fi

exec /bin/goalert --db-url "$DB_URL" --log-requests
