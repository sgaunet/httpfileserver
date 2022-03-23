FROM scratch
WORKDIR /
COPY httpfileserver /
COPY resources /
USER MyUser
CMD [ "/httpfileserver","-d","/mount" ]
