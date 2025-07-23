create table if not exists post (
  _created_at datetime not null,
  _updated_at datetime not null,
  _deleted_at datetime,

  id text primary key,
  title text not null,
  slug text not null,
  created_at datetime not null,
  updated_at datetime not null,
  revision integer not null,
  content text not null
); 
