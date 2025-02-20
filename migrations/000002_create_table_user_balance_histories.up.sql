CREATE TABLE public."user_balance_histories" (
    "id" SERIAL NOT NULL,
    "user_id" INT NOT NULL,
    "no_rekening" VARCHAR(50) NOT NULL,
    "previous_balance" BIGINT NOT NULL,
    "change_amount" BIGINT NOT NULL,
    "new_balance" BIGINT NOT NULL,
    "transaction_type" VARCHAR(20) NOT NULL,
    "transaction_id" VARCHAR(50) NOT NULL,
    "created_at" INT NOT NULL,     
    CONSTRAINT user_balance_histories_pkey PRIMARY KEY ("id"),
    CONSTRAINT user_balance_histories_user_fk FOREIGN KEY ("user_id") REFERENCES public."users"("id") ON DELETE CASCADE,
    CONSTRAINT user_balance_histories_no_rekening_fk FOREIGN KEY ("no_rekening") REFERENCES public."users"("no_rekening") ON DELETE CASCADE
);

CREATE INDEX user_balance_histories_user_idx ON public."user_balance_histories"("user_id");
CREATE INDEX user_balance_histories_no_rekening_idx ON public."user_balance_histories"("no_rekening");
