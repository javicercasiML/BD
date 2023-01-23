
-- 1. Agregar una película a la tabla movies.

insert into movies (created_at, updated_at, title, rating, awards, release_date, length, genre_id) values
	("2023-03-10", "2023-03-10", "Dark", 9.9, 10, "2023-03-10", 8, 6);

-- 2. Agregar un género a la tabla genres.
insert into genres (created_at, updated_at, name, ranking, active) values
	("2023-03-10", "2023-03-10", "Miedo", 13, 1);
    
select * from movies;

-- 3. Asociar a la película del punto 1. genre el género creado en el punto 2.

select * from movies;
UPDATE movies SET genre_id = 13 WHERE title = "Dark" and id = 22;


-- 4. Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el punto 1.

select * from actors;
UPDATE actors SET favorite_movie_id = 22 WHERE last_name = "Weaver" ;


-- 5. Crear una tabla temporal copia de la tabla movies.

create temporary table temporary_movies select * from movies;
# CREATE TEMPORARY TABLE temp_table AS (SELECT column1, column2 FROM original_table WHERE condition);
select * from movies;
select * from temporary_movies;

-- 6. Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.

DELETE from temporary_movies WHERE awards < 5;
# SET SQL_SAFE_UPDATES = 0;
# SET SQL_SAFE_UPDATES = 1; Colocar safe mode


-- 7. Obtener la lista de todos los géneros que tengan al menos una película.
select * from movies;

select g.name, COUNT(*) as total
FROM genres g JOIN movies m on
g.id = m.genre_id
GROUP BY g.name HAVING total >= 1;

-- 8. Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.

select a.first_name, a.last_name, a.favorite_movie_id, m.awards from actors a
join movies m on
a.favorite_movie_id = m.id
WHERE m.awards > 3
ORDER BY m.awards;

select * from movies;

-- 9. Crear un índice sobre el nombre en la tabla movies.

create index movies_title_index on movies_db.movies(title);

-- 10. Chequee que el índice fue creado correctamente.
SHOW INDEX from movies_db.movies;

-- 11. En la base de datos movies ¿Existiría una mejora notable al crear índices? Analizar y justificar la respuesta.

-- Existe una mejora notable al crear indices en la tabla movies, ya que si hacemos un "explain select"
-- Vemos que la cantidad de filas recorridas para el armado de la respuesta, es considerablemente menor.


-- 12. ¿En qué otra tabla crearía un índice y por qué? Justificar la respuesta.

-- Se podría crear un indice de autores por apellido, ya que son columnas utilizadas para consulta,
-- y no tanto para actualización. Éstos son requeridos constantemente.