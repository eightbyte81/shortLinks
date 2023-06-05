CREATE TABLE "default_link" (
    "id" serial NOT NULL,
    "link" TEXT NOT NULL,
    CONSTRAINT "default_link_pk" PRIMARY KEY ("id")
) WITH (
    OIDS=FALSE
);

CREATE TABLE "short_link" (
    "id" serial NOT NULL,
    "link" TEXT NOT NULL,
    CONSTRAINT "short_link_pk" PRIMARY KEY ("id")
) WITH (
    OIDS=FALSE
);

CREATE TABLE "link_chain" (
    "id" serial NOT NULL,
    "default_link_id" serial NOT NULL,
    "short_link_id" serial NOT NULL,
    CONSTRAINT "link_chain_pk" PRIMARY KEY ("id")
) WITH (
    OIDS=FALSE
);

ALTER TABLE "link_chain" ADD CONSTRAINT "link_chain_fk0" FOREIGN KEY ("default_link_id") REFERENCES "default_link"("id");
ALTER TABLE "link_chain" ADD CONSTRAINT "link_chain_fk1" FOREIGN KEY ("short_link_id") REFERENCES "short_link"("id");