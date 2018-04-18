#!/bin/sh

set -eux

user_name="snippetbox"
password="snippetbox"
db_name="snippetbox"

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOF
    CREATE USER $user_name WITH PASSWORD '$password';
    CREATE DATABASE $db_name;
    GRANT ALL PRIVILEGES ON DATABASE $db_name TO $user_name;
EOF

export PGPASS=$password

psql -v ON_ERROR_STOP=1 --username $user_name $db_name <<-EOF
    CREATE TABLE snippets (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        content TEXT NOT NULL,
        created TIMESTAMP WITH TIME ZONE NOT NULL,
        expires TIMESTAMP WITH TIME ZONE NOT NULL
    );

    CREATE INDEX snippets_expires ON snippets (expires);
EOF

set +u
if [ -z "$INSERT_SAMPLE_DATA" ] ; then
  exit 0
fi
set -u

psql -v ON_ERROR_STOP=1 --username $user_name $db_name <<-EOF
INSERT INTO snippets (title, content, created, expires) VALUES (
    'An old silent pond',
    'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP + INTERVAL '1 YEAR'
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'Over the wintry forest',
    'Over the wintry\nforest, winds howl in  rage\nwith no leaves to blow.\n\n– Natsume Soseki',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP + INTERVAL '1 YEAR'
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'First autumn morning',
    'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP + INTERVAL '1 MINUTE'
);
EOF
