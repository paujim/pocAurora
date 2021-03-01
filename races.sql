DROP TABLE Racing.Races;
CREATE TABLE Racing.Races
(
    race_id varchar(36) NOT NULL,
    race_name TEXT NOT NULL,
    race_number int NOT NULL,
    meeting_id varchar(36) NOT NULL,
    meeting_name TEXT NOT NULL,
    category_id varchar(36) NOT NULL,
    advertised_start VARCHAR(255) NOT NULL
);

INSERT INTO Racing.Races
    (
    race_id,
    race_name,
    race_number,
    meeting_id,
    meeting_name,
    category_id,
    advertised_start
    )
VALUES
    (
        '759d7dea-e763-4d41-9351-95da0f7fbac3',
        'Tab Download The App (Bm69)',
        2,
        'da20428b-bd21-412e-bea4-0a7b625a0778',
        'Townsville',
        '4a2788f8-e825-4d36-9894-efd4baf1cfae',
        '1579236600'
 );

INSERT INTO Racing.Races
    (
    race_id,
    race_name,
    race_number,
    meeting_id,
    meeting_name,
    category_id,
    advertised_start
    )
VALUES
    (
        '7d9aaf4e-556d-4ed6-932f-c3c5b77eb1ec',
        'Sizzle Here Feb 7 - "The Great Kiwi Bbq" Trot',
        1,
        'a959f9bf-2f49-4089-a8a1-ee0acd823886',
        'Alexandra Park',
        '161d9be2-e909-4326-8c2c-35ed71fb460b',
        '1579237140'
 );