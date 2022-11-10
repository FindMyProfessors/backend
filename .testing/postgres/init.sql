create table professors
(
    id         serial,
    school_id  integer not null,
    first_name varchar not null,
    last_name  varchar not null,
    rmp_id     varchar not null,
    constraint professors_pk
        primary key (id)
);

create table schools
(
    id   serial,
    name varchar not null,
    constraint schools_pk
        primary key (id)
);

create table reviews
(
    id           serial,
    quality      double precision not null,
    difficulty   double precision,
    time         timestamp        not null,
    tags         character varying[] not null,
    grade        varchar,
    professor_id integer,
    constraint reviews_pk
        primary key (id),
    constraint reviews_professors_null_fk
        foreign key (professor_id) references professors
);

create table courses
(
    id        serial,
    name      varchar not null,
    code      varchar not null,
    school_id integer not null,
    constraint courses_pk
        primary key (id),
    constraint courses_schools_null_fk
        foreign key (school_id) references schools
);

create table professor_courses
(
    professor_id integer not null,
    course_id    integer not null,
    constraint professor_courses_courses_id_fk
        foreign key (course_id) references courses,
    constraint professor_courses_professors_id_fk
        foreign key (professor_id) references professors
);
