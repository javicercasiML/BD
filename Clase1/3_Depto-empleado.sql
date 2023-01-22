drop database if exists deptos;
create database deptos;
use deptos;

create table depto (
	nro varchar(50) not null,
    nombre varchar(50),
    localidad varchar(50),
    primary key (nro)
);

create table empleado (
	cod_emp varchar(50) not null,
    nombre varchar(50),
    apellido varchar(50),
    puesto varchar(50),
    fecha_alta datetime,
    salario int,
    comision int,
    depto_nro varchar(50),
    primary key (cod_emp),
    foreign key (depto_nro) references depto(nro)
);

insert into depto (nro, nombre, localidad) values
	("D-000-1", "Software", "Los Tigres"),
    ("D-000-2", "Sistemas", "Guadalupe"),
    ("D-000-3", "Contabilidad", "La Roca"),
    ("D-000-4", "Ventas", "Plata")
;

insert into empleado (cod_emp, nombre, apellido, puesto, fecha_alta, salario, comision, depto_nro) values
	("E-0001", "César", "Piñero", "Vendedor", "2018-05-12", 80000, 15000, "D-000-4"),
    ("E-0002", "Yosep", "Kowaleski", "Analista", "2015-07-14", 140000, 0, "D-000-2"),
    ("E-0003", "Mariela", "Barrios", "Director", "2014-06-05", 185000, 0, "D-000-3"),
    ("E-0004", "Jonathan", "Aguilera", "Vendedor", "2015-06-03", 85000, 10000, "D-000-4"),
    ("E-0005", "Daniel", "Brezezicki", "Vendedor", "2018-03-03", 83000, 10000, "D-000-4"),
	("E-0006", "Mito", "Barchuk", "Presidente", "2014-06-05", 190000, 0, "D-000-3"),
    ("E-0007", "Emilio", "Galarza", "Desarrollador", "2014-08-02", 60000, 0, "D-000-1")
;

-- 1. Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores.
select e.nombre, e.apellido, e.puesto, d.localidad, e.puesto
from depto as d join empleado as e on
d.nro = e.depto_nro
WHERE e.puesto = "Vendedor";

-- 2. Visualizar los departamentos con más de cinco empleados.
select d.nombre, count(*) as Cantidad
from depto as d join empleado as e on
d.nro = e.depto_nro
GROUP BY d.nro having COUNT(d.nro) > 1;


-- 3. Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’.
select e.nombre, e.salario, d.nombre, e.puesto
from depto as d join empleado as e on
d.nro = e.depto_nro
WHERE e.puesto = 
(select empleado.puesto from empleado WHERE empleado.nombre = "Mito" and empleado.apellido = "Barchuk");

-- 4. Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre.
select e.*
from depto as d join empleado as e on
d.nro = e.depto_nro
WHERE d.nombre = "Contabilidad"
ORDER BY e.nombre; 

select * from empleado as e where e.depto_nro in (select nro from depto where nombre = "Contabilidad") order by nombre;


-- 5. Mostrar el nombre del empleado que tiene el salario más bajo.
select e.nombre, e.salario
from empleado as e 
WHERE e.salario IN (select min(empleado.salario) from empleado);

select nombre from empleado ORDER BY empleado.salario LIMIT 1;


-- 6. Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.
SELECT e. * FROM empleado as e
WHERE e.salario = (
	SELECT MAX(e2.salario) FROM empleado as e2 WHERE e2.depto_nro = (
		SELECT d.nro FROM depto as d WHERE d.nombre = "Ventas"));