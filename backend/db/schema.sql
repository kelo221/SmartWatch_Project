CREATE TABLE "user"
(
    id         serial primary key,
    "username" varchar not null,
    "password" varchar not null
);

CREATE TABLE "event"
(
    id         serial primary key,
    eventName  varchar not null,
    created_at timestamp default now(),
    eventId    int     not null,
    constraint fk_author_id foreign key (eventId) references "user" (id)
);