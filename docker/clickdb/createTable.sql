CREATE DATABASE Test;

CREATE TABLE if not exists Test.prj_log (
    loguui UUID NOT NULL,
    ip String NOT NULL,
    useruuid String NOT NULL,
    tstamp Int NOT NULL,
    logurl String NOT NULL,
    datarequest String NOT NULL ,
    dataresponse String NOT NULL
) 
ENGINE = TinyLog
