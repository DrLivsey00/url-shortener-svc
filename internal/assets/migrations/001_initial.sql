-- +migrate Up
create table links
    {
        long text not null,
        short text not null
    }
create index link_index on links (long);
-- +migrate Down
