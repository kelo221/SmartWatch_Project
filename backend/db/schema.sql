CREATE TABLE "user"
(
    id         serial primary key,
    "username" varchar not null UNIQUE,
    "password" varchar not null
);

CREATE TABLE "event"
(
    id             serial primary key,
    eventName      varchar                 not null,
    eventTime      timestamp               not null UNIQUE,
    created_at     timestamp default now(),
    userID         int                     not null,
    isSilent       bool      default false not null,
    snoozeDisabled bool      default false not null,
    constraint fk_event_id foreign key (userID) references "user" (id)
);