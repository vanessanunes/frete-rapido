create table dispatcher (
    id varchar primary key,
    request_id varchar NOT NULL,
    registered_number_shipper varchar NOT NULL,
    registered_number_dispatcher varchar NOT NULL,
    zipcode_origin int NOT NULL
);

create table offer (
    id serial primary key,
    id_dispatcher varchar NOT NULL,
    offer int NOT NULL,
    table_reference varchar NOT NULL,
    simulation_type int NOT NULL,
    service varchar NOT NULL,
    delivery_time_days int NOT NULL,
    delivery_time_estimated_date timestamp NOT NULL,
    expiration timestamp,
    cost_price numeric(5,2) NOT NULL,
    final_price numeric(5,2) NOT NULL,
    weights_real int NOT NULL,
    weights_used int NOT NULL,
    original_delivery_time_days int NOT NULL,
    original_delivery_time_estimated_date timestamp NOT NULL,
    foreign key (id_dispatcher) REFERENCES dispatcher(id)
);

create table carrier (
    id serial primary key,
    id_offer int,
    name varchar NOT NULL,
    registered_number varchar NOT NULL,
    state_inscription varchar NOT NULL,
    logo varchar NOT NULL,
    reference int NOT NULL,
    company_name varchar NOT NULL,
    foreign key (id_offer) REFERENCES offer (id)
);