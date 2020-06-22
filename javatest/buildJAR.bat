java HelloWorld

REM The 'e' flag (for 'entrypoint') creates or overrides the manifest's Main-Class attribute
jar cfe Hello.jar HelloWorld *.class
java -jar Hello.jar

REM add an entrypoint into the manifest from "Manifest.txt"
jar cfm World.jar Manifest.txt *.class
java -jar World.jar
