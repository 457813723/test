FROM ac2-registry.cn-hangzhou.cr.aliyuncs.com/ac2/base:ubuntu22.04-latest
WORKDIR /usr/local
COPY myapp .
EXPOSE 8899

ENTRYPOINT ["./myapp"]
