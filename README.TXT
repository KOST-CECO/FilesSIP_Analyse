Kurzanleitung
-------------
Das Programm FilesSIP_Analyse.exe ist ein Kommandozeilen Programm und muss unter Windows in einem CMD Fenster oder eine PowerShell Fenster ausgef�hrt werden.

Der Programmaufruf erfolgt mit Angabe eines Ordners als Parameter.
Dieser Ordner und alle Unterordner werden nach eCH-0160 und arelda SIPs durchsucht und in allen Files SIPs gepr�ft ob ein Verlinkung von Ordnungssystem auf Dateiablage mit Hilfe des <dateiRef> Tags vorliegt.
Das Programm schreibt f�r jeden abgearbeiteten Ordner ein Fortschrittszeichen "." auf die Error Ausgabe.
F�r fehlerhafte SIPs wird Pfad und SIP Namen ausgegeben.
Es werden auch gezippte SIPs untersucht.

Programmaufruf: FilesSIP_Analyse.exe Ordner

Mit diesem Aufruf werden die Fortschrittszeichen nicht angezeigt:
FilesSIP_Analyse.exe Ordner 2> null

Mit diesem Aufruf werden alle defekten SIP Name in eine Datei geschrieben:
FilesSIP_Analyse.exe Ordner > out.txt

Zusatz f�r das Schweizerische Bundesarchiv: 
Es werden auch ungezippte SIPs/AIPs erkannt wo sich metadata.xml hier befindet "header/old/SIP/metadata.xml"
