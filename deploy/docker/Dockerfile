FROM makeblock/alpine

ARG tmp_api_version=default_value
ENV API_VERSION=$tmp_api_version

ADD ./build /match-schedule/

WORKDIR /match-schedule

EXPOSE 8000

CMD ["./main"]