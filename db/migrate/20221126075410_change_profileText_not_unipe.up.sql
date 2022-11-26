--  プロフィールテキストのNULL制約を外す
ALTER TABLE "users" ALTER COLUMN "profile_text" DROP NOT NULL;

-- generate migrate file command
-- migrate create -ext sql -dir db/migrate [title]
