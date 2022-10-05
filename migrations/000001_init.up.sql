CREATE TABLE IF NOT EXISTS public."user" (
    id serial4 NOT NULL,
    "name" varchar(255) NOT NULL,
    username varchar(50) NOT NULL,
    email varchar(255) NOT NULL,
    "password" varchar(255) NOT NULL,
    status smallint NOT NULL DEFAULT 0,
    created_date timestamp(0) NOT NULL DEFAULT now(),
    updated_date timestamp(0) NULL,
    last_login_date timestamp(0) NULL,
    CONSTRAINT user_username_key UNIQUE (username),
    CONSTRAINT user_email_key UNIQUE (email),
    CONSTRAINT user_pkey PRIMARY KEY (id)
    );
CREATE INDEX user_username_idx ON public.user USING btree (username);
CREATE INDEX user_email_idx ON public.user USING btree (email);
COMMENT ON COLUMN "user"."status" IS '0 - disabled, 5 - waiting confirmation, 10 - active';

INSERT INTO public."user" ("name", username, email, "password", status) VALUES('admin', 'admin', 'admin@local.host', '$2a$08$HwTmFXKqGgsiisAphWBN8O5ZdCejEhILJik9IoFXSO1zUzCzi.IGW', 10);

CREATE TABLE IF NOT EXISTS public."auth_item" (
    "name" varchar(90) NOT NULL,
    "type" smallint NOT NULL DEFAULT 0,
    "parent" varchar(90) NULL,
    "description" text NULL,
    CONSTRAINT auth_item_pkey PRIMARY KEY ("name")
    );
CREATE INDEX auth_item_parent_idx ON public.auth_item USING btree (parent);
ALTER TABLE public."auth_item" ADD CONSTRAINT auth_item_parent_fk FOREIGN KEY (parent) REFERENCES public."auth_item"("name");
COMMENT ON COLUMN "auth_item"."type" IS '1 - role, 2 - action';

INSERT INTO "public".auth_item ("name", "type", "parent", "description") VALUES ('guest', 1, NULL, 'Неавторизованный');
INSERT INTO "public".auth_item ("name", "type", "parent", "description") VALUES ('user', 1, 'guest', 'Авторизованный пользователь');
INSERT INTO "public".auth_item ("name", "type", "parent", "description") VALUES ('manager', 1, 'user', 'Менеджер бекенда');
INSERT INTO "public".auth_item ("name", "type", "parent", "description") VALUES ('admin', 1, 'manager', 'Администратор');
INSERT INTO "public".auth_item ("name", "type", "parent", "description") VALUES ('User.Login', 2, 'guest', 'Авторизация');
INSERT INTO "public".auth_item ("name", "type", "parent", "description") VALUES ('User.Logout', 2, 'guest', 'Выйти');
INSERT INTO "public".auth_item ("name", "type", "parent", "description") VALUES ('User.Status', 2, 'guest', 'Статус пользователя');
INSERT INTO "public".auth_item ("name", "type", "parent", "description") VALUES ('User.Profile', 2, 'user', 'Профиль полльзователя');
INSERT INTO "public".auth_item ("name", "type", "parent", "description") VALUES ('Admin.Modules', 2, 'manager', 'Список модулей бекенда');

CREATE TABLE IF NOT EXISTS public."auth_item_assignment" (
    "username" varchar(50) NOT NULL,
    "item_name" varchar(90) NOT NULL,
    CONSTRAINT auth_item_assignment_pkey PRIMARY KEY ("username", "item_name")
    );
CREATE INDEX auth_item_assignment_user_idx ON public.auth_item_assignment USING btree (username);
ALTER TABLE public."auth_item_assignment" ADD CONSTRAINT auth_item_assignment_user_fk FOREIGN KEY (username) REFERENCES public."user"(username);
ALTER TABLE public."auth_item_assignment" ADD CONSTRAINT auth_item_assignment_item_fk FOREIGN KEY (item_name) REFERENCES public."auth_item"("name");

INSERT INTO public.auth_item_assignment ("username", "item_name") VALUES ('admin', 'admin');

CREATE TABLE IF NOT EXISTS public."session" (
    id uuid NOT NULL,
    "key" varchar(100) NOT NULL,
    user_id int4 NULL,
    created_date timestamp(0) NOT NULL DEFAULT now(),
    updated_date timestamp(0) NULL,
    allow_auto_login bool NOT NULL DEFAULT false,
    "data" jsonb NULL,
    "access" jsonb NULL,
    "ip" varchar(100) NOT NULL,
    "meta" jsonb NULL,
    CONSTRAINT session_key_key UNIQUE ("key"),
    CONSTRAINT session_pkey PRIMARY KEY (id)
    );
CREATE INDEX session_user_idx ON public.session USING btree (user_id);
CREATE INDEX session_key_idx ON public.session USING btree ("key");
CREATE INDEX session_updated_date_idx ON public.session USING btree (updated_date);
ALTER TABLE public."session" ADD CONSTRAINT session_user_fk FOREIGN KEY (user_id) REFERENCES public."user"(id);