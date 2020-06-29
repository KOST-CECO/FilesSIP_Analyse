@ECHO OFF
SETLOCAL

SET GNUDIR=C:\Software\PCUnixUtils

IF [%2]==[] (
        ECHO usage: %0 XSD XML
        EXIT /B
)

IF NOT EXIST %1 (
        ECHO XSD file "%1" not found
        ECHO usage: %0 XSD XML
        EXIT /B
)

IF NOT EXIST %2 (
        ECHO XML file "%2" not found
        ECHO usage: %0 XSD XML
        EXIT /B
)

timeit.exe -f timeit.dat -k xml %GNUDIR%\xmllint.exe --noout --stream --schema "%1" "%2"

EXIT /B
