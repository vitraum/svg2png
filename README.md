
docker-compose up --build --remove-orphans

curl -d @test.svg http://localhost:8001/api/v1/proxy/namespaces/admin/services/svg2png:80/v1/png > test.png
