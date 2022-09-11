CREATE TABLE prj_log (
    loguui CHARACTER VARYING(32) NOT NULL UNIQUE,
    ip CHARACTER VARYING(32) NOT NULL,
    useruuid CHARACTER VARYING(32) NOT NULL UNIQUE,
    tstamp CHARACTER VARYING(32) NOT NULL UNIQUE,
    logurl CHARACTER VARYING(32) NOT NULL UNIQUE,
    datarequest CHARACTER VARYING(32) NOT NULL UNIQUE,
    dataresponse CHARACTER VARYING(32) NOT NULL UNIQUE,
);

	-- const query = `INSERT INTO prj_log(loguui, ip, useruuid, timestamp, url, datarequest, dataresponse) VALUES (:loguui, :ip, :useruuid, :timestamp, :url, :datarequest, :dataresponse) ON CONFLICT DO NOTHING;`

