FROM busybox

COPY ./simpleserv /simpleserv

EXPOSE 8196

CMD ["/simpleserv"]
