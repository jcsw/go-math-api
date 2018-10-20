# Go Math API

### Math API writen in Golang ###
#### Modeling based on 'clean architecture' ####

Required
```
go 1.9
docker-compose
```

Build
```
make deps
make all
```

Dependencies to run
```
cd docker/mongodb
docker-compose up
```

Run
```
make run-deps
make run
```

### Stress testing with Locust ###

To install with Ubuntu
```
sudo apt-get install python3 python3-pip
python3 -m pip install locustio
```

To Run
```
locust -f stress-test/math-api-stress-test.py --host=http://localhost:9900 --csv=stress-test/report/math-api
```

To access web ui
```
http://localhost:8089/
```

To Run without the web UI
```
locust -f stress-test/math-api-stress-test.py --host=http://localhost:9900 --no-web -c 100 -r 10 --csv=stress-test/report/math-api
```

Locust Documentation
```
https://docs.locust.io/en/latest/what-is-locust.html
```