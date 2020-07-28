FROM alpine:latest
LABEL maintainer="liion"
# 拷贝编译程序
ADD main /
RUN chmod +x /main
WORKDIR /
# 打开8080端口
EXPOSE 8080
# 运行!
CMD ["./main"]