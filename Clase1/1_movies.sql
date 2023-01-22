# 1
select * from movies;

# 2
select first_name, last_name, rating from actors;

# 3
select title as titulo from series as series;

# 4
select first_name, last_name, rating from actors where rating > 7.5;

# 5
select title,rating,awards from movies where rating > 7.5 and awards > 2;

# 6
select title, rating from movies order by rating;

# 7
select id,title as titulo from movies limit 3;

# 8
select title,rating from movies order by rating desc limit 5;

# 9
select first_name, last_name, rating from actors limit 10;

# 10
select title, rating from movies where title like "Toy Story%";

# 11
select first_name, last_name, rating from actors where first_name like "sam%";

# 12
select title as Titulo, release_date as Lanzamiento from movies where release_date BETWEEN '2004-01-01' and '2008-12-12'; 

SELECT title,YEAR(release_date) FROM movies WHERE YEAR(release_date) >= 2004 AND YEAR(release_date) <= 2008;

# 13
select title,rating,awards,release_date from movies 
where rating > 3 and awards > 2 and release_date BETWEEN '1998-01-01' and '2009-12-12'
order by rating; 

# 14
SELECT *
FROM actor_movie
WHERE movie_id IN (SELECT id FROM movies WHERE rating=9.0);