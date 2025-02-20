CREATE TABLE public."user"
(
    "id" SERIAL NOT NULL,
    "nama" VARCHAR(100) NOT NULL,
    "nik" VARCHAR(32) NOT NULL UNIQUE,
    "no_hp" VARCHAR(15) NOT NULL UNIQUE,
    "created_at" INT NOT NULL,
    "updated_at" INT NOT NULL,
    "deleted_at" INT,
    CONSTRAINT user_pkey PRIMARY KEY ("id")
);

CREATE INDEX user_nik_idx ON public."user"("nik");
CREATE INDEX user_no_hp_idx ON public."user"("no_hp");
CREATE INDEX user_deleted_at_idx ON public."user"("deleted_at");

