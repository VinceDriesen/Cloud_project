Dit is mijn tweede poging voor het school project van Cloud computing

-- Laravel deel --
Hierin komt het volledige laravel project waarin in alle api's ga willen aansturen.

Eerst de docker file runnen : 
    docker-compose up - d

Ik maak hier gebruik van de database 'maria-db' zoals ook in de les is gebruikt

Om in de docker/shell te gaan maak dan gebruik van 
    docker exec -it hospital_laravel /bin/bash

Hierna zit je dan in de folder app, waarna je in de hospitalWebsite moet gaan. Hier staat nu je laravel project
Nu kan je ook alle commands van php artisan gebruiken. Zolang je in de docker shell blijft

Runnen kan met:
    php artisan serve --host 0.0.0.0
    npm run dev of npm run build

Indien pagina te lang doet over refreshen volgende commandos:
    php artisan cache:clear
    php artisan config:clear
    php artisan route:clear
    php artisan view:clear

Om je laravel database te zien in php storm moet je een MariaDB connection maken. Dit kan je doen door een inspect om ze de host er uit te krijgen:
    docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' hospital_laravel_db
Deze zou dan iets in de respons geven van:
    172.18.0.3
Nu kan je verbinden als:
    Host: 172.18.0.3
    port: 3306
    user: root
    password: root
    database: hospital_database
