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

CREATE TABLE "default_short_links" (
    "id" serial NOT NULL,
    "default_link_id" serial NOT NULL,
    "short_link_id" serial NOT NULL,
    CONSTRAINT "default_short_links_pk" PRIMARY KEY ("id")
) WITH (
    OIDS=FALSE
);

ALTER TABLE "default_short_links" ADD CONSTRAINT "default_short_links_fk0" FOREIGN KEY ("default_link_id") REFERENCES "default_link"("id");
ALTER TABLE "default_short_links" ADD CONSTRAINT "default_short_links_fk1" FOREIGN KEY ("short_link_id") REFERENCES "short_link"("id");