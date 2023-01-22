USE movies_db;
# 1:
select series.title as titulo, genres.name as Genero
from series inner join genres ON
series.genre_id = genres.id;

# 2:
select e.title, a.first_name, a.last_name
from episodes as e inner join actor_episode as ae on
e.id = ae.episode_id 
INNER JOIN actors as a on
a.id = ae.actor_id;

# 3:
select s.title as Series, COUNT(t.serie_id) as Temporadas
from series as s join seasons as t on
t.serie_id = s.id
GROUP BY s.title;

# 4:
SELECT g.name as Genero, COUNT(m.genre_id) as Peliculas
FROM genres g JOIN movies m ON
m.genre_id = g.id
GROUP BY g.name HAVING Peliculas >= 3;

# 5_ Mostrar sólo el nombre y apellido de los actores que trabajan en 
# todas las películas de la guerra de las galaxias y que estos no se repitan.
SELECT DISTINCT a.first_name, a.last_name
FROM actors as a WHERE a.id in
(SELECT am.actor_id 
 from movies as m join actor_movie as am on
 am.movie_id = m.id
 WHERE m.title LIKE '%La Guerra de las galaxias%');
