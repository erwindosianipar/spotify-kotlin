create database kotlify;

use kotlify;

create table artist (
    id int not null auto_increment,
    name varchar(20),
    debut varchar(20),
    category varchar(20),
    image varchar(255),
    primary key(id)
);

create table genre (
    id int not null auto_increment,
    name varchar(20),
    primary key(id)
);

create table song (
    id  int not null auto_increment,
    genre_id int,
    artist_id int,
    name varchar(20),
    primary key(id),
    foreign key(genre_id) references genre(id),
    foreign key(artist_id) references artist(id)
);

insert into artist (name, debut, category, image) values
('Rich Brian', '2018-03-21', 'Solo', 'https://www.provoke-online.com/images/spotify_on_stage_jkt_2019_-_cover_playlist.jpg'),
('Jonas Blue', '2015-11-11', 'Solo', 'https://i.pinimg.com/originals/cd/5e/29/cd5e2992d42c61be4e767191cd5b8d86.jpg'),
('Tulus', '2016-03-12', 'Solo', 'https://pbs.twimg.com/media/DlbyavcVsAAmbIo.jpg'),
('Nidji', '2015-10-12', 'Group', 'https://i.pinimg.com/originals/b5/f6/84/b5f68482dc5f481bcfa89c2f34097e52.png');

insert into genre (name) values
('Pop'),
('Jazz'),
('Country'),
('Rock'),
('Blues');

insert into song (genre_id, artist_id, name) values
(1, 1, '100 Degres'),
(3, 1, 'Drive safe'),
(4, 1, 'Chaos'),
(5, 2, 'Mama'),
(3, 2, 'I See Love'),
(4, 3, 'Sepatu'),
(2, 3, 'Gajah'),
(5, 3, 'Teman Hidup'),
(5, 4, 'Laskar Pelangi'),
(2, 4, 'Disco Lazy Time'),
(1, 4, 'Diatas awan');

select * from artist;
select * from genre;
select * from song;
