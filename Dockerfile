FROM golang:alpine

RUN sed -i 's/http:\/\/dl-cdn.alpinelinux.org/https:\/\/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && \
  apk update && apk upgrade && \
  apk add mdocml-apropos coreutils gcc g++ libc-dev ca-certificates autoconf automake file libtool \
  bc tree vim git fish dialog less make tzdata gsl-dev

RUN sed -i "s/bin\/ash/usr\/bin\/fish/" /etc/passwd

RUN echo "set mouse-=a" >> ~/.vimrc

RUN mkdir -p /root/.config/fish && \
  # echo "set -gx GOPATH /data/gocode" >> /root/.config/fish/config.fish && \
  echo "set -gx GO111MODULE on" >> /root/.config/fish/config.fish

RUN rm -f /etc/localtime && ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

ADD . /root/go/src/github.com/tosone/gsl/

WORKDIR /root/go/src/github.com/tosone/gsl/

ENV SHELL /usr/bin/fish

CMD ["/usr/bin/fish"]
