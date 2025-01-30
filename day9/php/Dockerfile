FROM ubuntu:latest
RUN apt update && apt install -y php8.3
WORKDIR /root/
COPY main.php /root/
CMD ["php", "/root/main.php"]
