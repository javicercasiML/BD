drop database if exists biblioteca;
create database biblioteca;
use biblioteca;

create table libro (
	id int not null auto_increment,
    titulo varchar(100),
    editorial varchar(100),
    area varchar(100),
    primary key (id)
);

create table autor(
	id int not NULL AUTO_INCREMENT,
    nombre VARCHAR(100),
    nacionalidad VARCHAR(100),
    PRIMARY KEY (id)
);

create table libro_autor (
	autor_id int,
    libro_id int,
    foreign key (autor_id) references autor(id),
    foreign key (libro_id) references libro(id)
);

create table estudiante (
	id int not null auto_increment,
    nombre varchar(100),
    apellido varchar(100),
    direccion varchar(100),
    carrera varchar(100),
    edad int,
    primary key (id)
);

create table prestamo (
	estudiante_id int,
    libro_id int,
    fecha_prestamo timestamp,
    fecha_devolucion timestamp,
    devuelto bool,
    foreign key (estudiante_id) references estudiante(id),
    foreign key (libro_id) references libro(id)
);

insert into libro (titulo, editorial, area) values
	("Lengua 1", "Salamandra", "lengua"),
	("fisica 2", "Javi", "fisica"),
	("matematica 1", "Kapeluz", "matematica"),
	("El Universo: Guía de viaje", "Kapeluz", "matematica"),
	("Base de datos", "Kapeluz", "Internet")
;

insert into autor (nombre, nacionalidad) values
	("Messi", "Argentina"),
    ("Paredes", "Argentina"),
    ("Dibu", "Francia"),
    ("J.K. Rowling", "Argentina"),
    ("J.K. Rowling", "Italia")
;

insert into libro_autor (autor_id, libro_id) values
	(3, 2),
    (1, 5),
    (1, 4),
    (2, 2),
    (3, 5),
    (5, 4),
    (5, 1)
;

insert into estudiante (nombre, apellido, direccion, carrera, edad) values
	("Javier", "Cercasi", "Avenida 123", "Computacion", 24),
    ("Juan", "Gomez", "Avenida 123", "Computacion", 21),
    ("Filippo", "Galli", "Avenida 333", "Informatica", 2),
    ("Juana", "Viale", "Avenida 486", "Electronica", 25),
    ("Martina", "Sanchez", "Avenida 684", "Derecho", 20)
;

insert into prestamo (estudiante_id, libro_id, fecha_prestamo, fecha_devolucion, devuelto) values
	(1, 5, "2005-01-10", "2005-01-24", true),
    (2, 4, "2018-01-10", "2018-01-24", true),
    (3, 3, "2021-05-10", "2021-07-24", true),
    (4, 2, "2022-11-10", null, false),
    (5, 1, "2024-10-10", null, false),
    (4, 3, "2023-03-10", "2023-03-24", true),
    (3, 5, "2025-05-10", "2025-05-24", true),
    (2, 3, "2019-08-10", "2019-08-24", true)
;

# 1. Listar los datos de los autores.
select * from autor;

-- 2. Listar nombre y edad de los estudiantes.
select nombre, edad from estudiante;


-- 3. ¿Qué estudiantes pertenecen a la carrera informática?
select nombre, apellido, carrera from estudiante WHERE carrera = "Informatica";


-- 4. ¿Qué autores son de nacionalidad francesa o italiana?
select a.nombre, a.nacionalidad from autor a WHERE a.nacionalidad = "Francia" or a.nacionalidad = "Italia";


-- 5. ¿Qué libros no son del área de internet?
select titulo, editorial, area from libro WHERE area != "Internet";


-- 6. Listar los libros de la editorial Salamandra.
select id, titulo, area from libro where editorial = "Salamandra";

-- 7. Listar los datos de los estudiantes cuya edad es mayor al promedio.
select e.id, e.nombre, e.apellido, e.edad, e.carrera from estudiante e 
WHERE e.edad > (select AVG(estudiante.edad) from estudiante);


-- 8. Listar los nombres de los estudiantes cuyo apellido comience con la letra G.
select nombre, apellido FROM estudiante WHERE apellido like "G%";



-- 9. Listar los autores del libro “El Universo: Guía de viaje”. (Se debe listar solamente los nombres).

SELECT a.nombre from autor a 
WHERE a.id in (select lb.autor_id from libro_autor lb join libro l on 
lb.libro_id = l.id WHERE l.titulo LIKE "El Universo: Guía de viaje");


-- 10. ¿Qué libros se prestaron al lector “Filippo Galli”?
SELECT l.titulo from libro l 
WHERE l.id in (select p.libro_id from prestamo p join estudiante e on 
p.estudiante_id = e.id WHERE e.nombre = "Filippo" and e.apellido = "Galli");


-- 11. Listar el nombre del estudiante de menor edad.
select e.nombre from estudiante e 
WHERE e.edad = (select MIN(estudiante.edad) from estudiante);


-- 12. Listar nombres de los estudiantes a los que se prestaron libros de Base de Datos.
SELECT e.nombre from estudiante e 
WHERE e.id in (select p.estudiante_id from prestamo p join libro l on 
p.libro_id = l.id WHERE l.titulo LIKE "Base%");


-- 13. Listar los libros que pertenecen a la autora J.K. Rowling.
SELECT l.titulo from libro l 
WHERE l.id in (select la.libro_id from libro_autor la join autor a on 
la.autor_id = a.id WHERE a.nombre = "J.K. Rowling");


-- 14. Listar títulos de los libros que debían devolverse el 16/07/2021.
select l.titulo, p.fecha_devolucion from libro l join prestamo p on 
l.id = p.libro_id where p.fecha_devolucion = "2021-07-24";