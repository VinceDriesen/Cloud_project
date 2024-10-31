Dit is mijn tweede poging voor het school project van Cloud computing

-- Laravel deel --
Hierin komt het volledige laravel project waarin in alle api's ga willen aansturen.

Eerst de docker file runnen : 
    docker-compose up --build - d

Ik maak hier gebruik van de database 'maria-db' zoals ook in de les is gebruikt

Om in de docker/shell te gaan maak dan gebruik van 
    docker exec -it hospital_laravel /bin/bash

Hierna zit je dan in de folder app, waarna je in de hospitalWebsite moet gaan. Hier staat nu je laravel project
Nu kan je ook alle commands van php artisan gebruiken. Zolang je in de docker shell blijft

Runnen kan met, maar dit wordt echter door de dockerfile gedaan:
    php artisan serve --host 0.0.0.0
    npm run dev of npm run build

Indien pagina te lang doet over refreshen volgende commandos:
    php artisan cache:clear
    php artisan config:clear
    php artisan route:clear
    php artisan view:clear

Je kan met de database verbinden via adminer:
    - Systeem:  MySql
    - Server: db_laravel
    - Username: laravel
    - Password: laravelPwd
    - Database: hospital_laravel_database

-- Services - userProfileAPISrvice --
Deze service is voor het volledig profiel van de gebruiker op te vragen
Ik maak hier gebruik van postgresql als database en .NET SOAP in C# voor de service implementatie
Hiervoor heb ik gekozen omdat SOAP vaak in C# wordt geschreven, maar het ook zou kunnen in java. 
Echter, aangezien ik java al beter ken dan C# én java wordt nog bij andere services gebruikt heb ik gekozen voor C#

Om in de docker/shell te gaan maak dan gebruik van 
    docker exec -it user_profile_api /bin/bash

Na docker-compose up heb je vaak ook nog dotnet ef nodig voor de database, Dit kan je dan doen met volgende commandos:
    dotnet tool install --global dotnet-ef
    export PATH="$PATH:/app/.dotnet/tools"
Hierna kan je gebruik maken van:
    dotnet ef migration add ...
    dotnet ef database update

Je kan met de database verbinden via adminer:
    - Systeem:  MySql
    - Server: db_user_profile
    - Username: postgres
    - Password: postgres
    - Database: user_profile_database
