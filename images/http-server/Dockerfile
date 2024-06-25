FROM python:3-alpine

RUN mkdir /wwwroot
VOLUME /wwwroot

EXPOSE 8080

WORKDIR /wwwroot

CMD ["python", "-m" , "http.server", "8080"]
