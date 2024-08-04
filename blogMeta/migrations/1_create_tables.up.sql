create table BlogMeta
(
    slug          varchar(255) primary key,
    createdAt     timestamp    not null default current_timestamp,
    updatedAt     timestamp    not null default current_timestamp,
    title         varchar(255) not null,
    description   text         not null,
    featuredImage varchar(255) not null,
    tags          text[],
    author        varchar(255) not null
);

create index idx_blog_meta_updated_at on BlogMeta (updatedAt);

CREATE
OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updatedAt
= NOW();
RETURN NEW;
END;
$$
language 'plpgsql';

CREATE TRIGGER update_blog_meta_updated_at
    BEFORE UPDATE
    ON BlogMeta
    FOR EACH ROW
    EXECUTE PROCEDURE update_timestamp();