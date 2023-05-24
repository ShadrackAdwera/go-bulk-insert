-- name: CreateCase :one
INSERT INTO ecde_cases (date_rep,day,month,year,cases,deaths,countries_and_territories,geo_id,country_territory_code,continent_exp,load_date, iso_country)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) RETURNING *;

-- name: ListCases :many
SELECT * FROM ecde_cases
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetCase :one
SELECT * FROM ecde_cases
WHERE id = $1 LIMIT 1;

