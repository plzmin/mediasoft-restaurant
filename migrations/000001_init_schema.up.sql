create table if not exists menu
(
    uuid              uuid not null
        constraint menu_pk
            primary key,
    on_date           timestamp,
    opening_record_at timestamp,
    closing_record_at timestamp,
    created_at        timestamp default current_timestamp
);

create table if not exists products
(
    uuid        uuid not null
        constraint product_pk
            primary key,
    name        text,
    description text,
    type        integer,
    weight      integer,
    price       numeric,
    created_at  timestamp default current_timestamp
);

create table if not exists menu_product
(
    menu_uuid    uuid
        constraint menu_product_menu_uuid_fk
            references menu
            on update cascade on delete cascade,
    product_uuid uuid
        constraint menu_product_products_uuid_fk
            references products
            on update cascade on delete cascade
);

create table if not exists orders
(
    uuid       uuid not null
        constraint order_pk
            primary key,
    user_uuid  uuid,
    created_at timestamp default current_timestamp
);

create table if not exists order_item
(
    order_uuid   uuid
        constraint order_item_order_uuid_fk
            references orders
            on update cascade on delete cascade,
    count        integer,
    product_uuid uuid
);