create table "users"(id uuid primary key, first_name text not null, last_name text not null, email text not null, password text not null, created_at time, updated_at time, deleted_at time);

create table "products"( id uuid primary key, name text not null, cost float not null);

insert into products values ('87eeac1d-d391-4017-aa17-3795bcf9eb2f', 'El periódico', 14.99);