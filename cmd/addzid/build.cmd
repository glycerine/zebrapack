@echo off
echo>gitcommit.go package main
echo>>gitcommit.go var LASTGITCOMMITHASH string

REM Get GIT version
SET GIT_REV_PRE=func init() { LASTGITCOMMITHASH = "
SET GIT_REV_POST= "}
git rev-parse HEAD > _rev.txt
set /P GITREVISION=< _rev.txt
for %%a in (%GITREVISION%) do (
    echo/|set /p ="%GIT_REV_PRE%"
    echo/|set /p ="%%a"
    echo/|set /p ="%GIT_REV_POST%"
) >> gitcommit.go
del _rev.txt

go build && go install
