FROM alpine
RUN set -ex \
	&& apk add --no-cache ca-certificates
COPY list_instances /
#ENTRYPOINT ["./list_instances"]
CMD /list_instances
