ALTER TABLE "users" ADD constraint users_name_key unique("name");

-- generate migrate file command
-- migrate create -ext sql -dir db/migrate [title]
