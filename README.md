
docker-compose up --build --remove-orphans

# TEST

curl -d @test.svg http://localhost:8544/v1/png > test.png

# TODO

Split chrome runners into seperate pods to enable autoscaling
 -> needs support for a resizable, remote pool, not included in chromedp/chromedp
