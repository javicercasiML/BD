USE movies_db;


-- ------------------------Primera parte:
/*
1_ Un join es utilizado para combinar 2 o mas tablas entre sí,
retornando la intersección de ambas tablas mediante una condicion en común.

2_ Join o inner Join: Trae los datos de intersección entre 2 o mas tablas.

   Left Join: Trae todos los datos de la tabla izquierda, y los relacionados de
la segunda.

3_ Group By: Utilizado para agrupar datos por una columna pasada indicada como
argumento. Se utiliza en conjunto con las funciones de agregacion, para
agrupar los datos que se necesiten.

4_ Having: Similar al select, pero es utilizado para incluir condiciones 
en el resultado dado de un Group By. Afecta a un conjunto de registros.

5_ Tabla temporal: Sería utilizada cuando se necesita consultar datos 
	que son necesitados con frecuencia, y asi evitar utilizar multiples
    joins. Tambien la utilizaría para hacer pruebas. 
    
    En casos donde el resultado de una consulta sea accesorio, y se 
    utiliza para operaciones conretas, que finalizan cuando se cierra la
    sesión.
    Ej: Carrito de compras, una tabla temporal. Tambien para facturaciones.
		Para procesar los datos antes de enviarlos a una tabla temporal.

6_ Indices: Sería utilizada cuando necesito realizar consultar mas veloces, ya que
	son un puntero a una columna específica. También cuando una columna es
    consultada con frecuencia.
    
    No lo utilizaría siempre, ya que la existencia de los indices tienen un costo
    a nivel de recursos, y cuando los datos son actualizados con 
    frecuencia, consumiríamos mas recursos de los necesitados en el sistema,
    ya que se necesitarán actualizar también los índices.
    
    (Se almacenan en Betree o Hash)

7_ SELECT * from table_a JOIN table_b ON table_a.campo = table_b.campo
   SELECT * from table_a LEFT JOIN table_b ON table_a.campo = table_b.campo

*/

-- ------------------------Segunda parte:

# 1. Mostrar el título y el nombre del género de todas las series.
select series.title as titulo, genres.name as Genero
from series inner join genres ON
series.genre_id = genres.id;

# 2. Mostrar el título de los episodios, el nombre y apellido de los actores 
#    que trabajan en cada uno de ellos.
select e.title, a.first_name, a.last_name
from episodes e inner join actor_episode ae on
e.id = ae.episode_id 
INNER JOIN actors as a on
a.id = ae.actor_id;

# 3. Mostrar el título de todas las series y el total de temporadas que tiene 
#    cada una de ellas.
select s.title as Series, COUNT(t.serie_id) as Temporadas
from series as s join seasons as t on
t.serie_id = s.id
GROUP BY s.title;

# 4. Mostrar el nombre de todos los géneros y la cantidad total de películas 
#	 por cada uno, siempre que sea mayor o igual a 3.
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
 WHERE m.title LIKE 'La Guerra de las galaxias%');

-- Si se quiere hacer mediante un indice de title, solo se debe usar el comodin al final %.
 
 
SELECT DISTINCT a.first_name, a.last_name FROM actors a 
JOIN actor_movie am ON
a.id = am.actor_id
JOIN movies m ON 
m.id = am.movie_id
WHERE m.title LIKE '%La Guerra de las galaxias%';
 
