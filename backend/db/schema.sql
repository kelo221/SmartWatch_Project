CREATE TABLE "user"
(
    id         serial primary key,
    "username" varchar not null UNIQUE,
    "password" varchar not null
);

CREATE TABLE "event"
(
    id         serial primary key,
    eventName  varchar not null,
    created_at timestamp default now(),
    eventId    int     not null,
    isSilent   bool      default false,
    constraint fk_event_id foreign key (eventId) references "user" (id)
);