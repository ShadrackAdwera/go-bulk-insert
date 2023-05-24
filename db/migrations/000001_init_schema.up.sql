CREATE TABLE "ecde_cases" (
  "id" bigserial PRIMARY KEY,
  "date_rep" varchar NOT NULL,
  "day" int NOT NULL,
  "month" int NOT NULL,
  "year" int NOT NULL,
  "cases" bigint NOT NULL,
  "deaths" bigint NOT NULL,
  "countries_and_territories" varchar NOT NULL,
  "geo_id" varchar NOT NULL,
  "country_territory_code" varchar NOT NULL,
  "continent_exp" varchar NOT NULL,
  "load_date" varchar NOT NULL,
  "iso_country" varchar NOT NULL
);
