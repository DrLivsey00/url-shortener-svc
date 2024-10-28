-- +migrate Up
create table links (
    id integer primary key autoincrement,
    alias text not null unique,
    url text not null
);
create index link_index on links (alias);

-- +migrate Down
drop table links;