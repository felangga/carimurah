CREATE TABLE public.products (
    id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL,
    name character varying NOT NULL,
    original_price numeric DEFAULT 0,
    price numeric DEFAULT 0 NOT NULL,
    rating double precision DEFAULT 0,
    rating_average double precision DEFAULT 0,
    url text,
    url_img text,
    ext_id varchar(255) NOT NULL
    UNIQUE(ext_id)
);
