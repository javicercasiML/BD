
-- 1:
CREATE TEMPORARY TABLE movies_db.TWD (

id INT,
title VARCHAR(45),
release_date DATETIME NOT NULL,
season_id INT,
rating DECIMAL(3,1),
PRIMARY KEY(id)
);

INSERT INTO movies_db.TWD(id, title, release_date, season_id, rating)

SELECT movies_db.episodes.id, movies_db.episodes.title, movies_db.episodes.release_date, movies_db.episodes.season_id, movies_db.episodes.rating
FROM movies_db.episodes
JOIN movies_db.seasons ON 
movies_db.episodes.season_id = movies_db.seasons.id
JOIN movies_db.series ON 
movies_db.seasons.serie_id = movies_db.series.id
WHERE movies_db.series.id = 3;


SELECT *
FROM movies_db.TWD
WHERE season_id = 20;

-- 2:
create index actors_last_name_index on movies_db.actors(last_name);
SHOW INDEX from movies_db.actors;

-- Sin el index, la siguiente consulta recorreria 49 rows, en la siguiente es 1:
EXPLAIN SELECT *
from movies_db.actors
WHERE movies_db.actors.last_name = "Ford";
