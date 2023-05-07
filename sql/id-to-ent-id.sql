DELIMITER $$
BEGIN
  DECLARE _count INT;

  SELECT
    count(column_type)
  INTO
    _count
  FROM
    information_schema.columns
  WHERE
    table_schema = 'service_template'
    AND
    table_name = 'details'
    AND
    column_name = 'ent_id';

  IF _count <= 0 THEN
    ALTER TABLE details CHANGE id ent_id CHAR(36);
  END IF;

  SELECT
    count(column_type)
  INTO
    _count
  FROM
    information_schema.columns
  WHERE
    table_schema = 'service_template'
    AND
    table_name = 'details'
    AND
    column_name = 'id';

  IF _count <= 0; THEN
    ALTER TABLE details ADD COLUMN id INT UNSIGNED NOT NULL UNIQUE AUTO_INCREMENT;
    ALTER TABLE details DROP PRIMARY KEY;
    ALTER TABLE details ADD PRIMARY KEY (id);
  END IF;

END$$
