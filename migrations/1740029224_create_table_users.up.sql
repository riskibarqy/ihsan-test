CREATE TABLE public."users"
(
    "id" SERIAL NOT NULL,
    "nama" VARCHAR(100) NOT NULL,
    "nik" VARCHAR(16) NOT NULL UNIQUE,
    "no_hp" VARCHAR(15) NOT NULL UNIQUE,
    "no_rekening" VARCHAR(50) NOT NULL UNIQUE,
    "created_at" INT NOT NULL,
    "updated_at" INT NOT NULL,
    "deleted_at" INT,
    UNIQUE("no_hp","nik"),
    CONSTRAINT user_pkey PRIMARY KEY ("id")
);

CREATE INDEX user_nik_idx ON public."users"("nik");
CREATE INDEX user_no_hp_idx ON public."users"("no_hp");
CREATE INDEX user_no_rekening_idx ON public."users"("no_rekening");
CREATE INDEX user_deleted_at_idx ON public."users"("deleted_at");

