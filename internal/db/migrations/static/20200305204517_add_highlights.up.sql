BEGIN;

CREATE TABLE highlights (
    id bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    created_at timestamptz DEFAULT NOW() NOT NULL,

    channel_id bigint REFERENCES channels (id) NOT NULL,

    highlighted_at timestamptz NOT NULL, -- Distinct from created_at in order to allow for highlights in the past.
    started_at timestamptz,
    status text NOT NULL,
    game text NOT NULL
);

COMMIT;
