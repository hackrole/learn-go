
-- +migrate Up
create unique index people_unique_id_idx  on people(id);

-- +migrate Down
drop index people_unique_id_idx;
