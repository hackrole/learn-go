create table "user" (
    user_id char(64) primary key,
    name text not null,
    credit bigint default 0
);

create table hash_text (
    hash char(64) primary key
    text TEXT
)
