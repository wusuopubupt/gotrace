FROM buildpack-deps:jessie-scm

# gcc for cgo
RUN apt-get update && apt-get install -y --no-install-recommends \
		g++ \
		gcc \
		libc6-dev \
		make \
		patch \
	&& rm -rf /var/lib/apt/lists/*

ENV GOLANG_VERSION 1.6.2
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 e40c36ae71756198478624ed1bb4ce17597b3c19d243f3f0899bb5740d56212a

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
	&& echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - \
	&& mkdir -p /usr/local/gotemp && tar -C /usr/local/gotemp -xzf golang.tar.gz \
	&& rm golang.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/gotemp/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

ENV GOLANG_TO_BUILD_URL https://golang.org/dl/go$GOLANG_VERSION.src.tar.gz
RUN curl -fsSL "$GOLANG_TO_BUILD_URL" -o golang.tar.gz
RUN	tar -C /usr/local -xzf golang.tar.gz && rm golang.tar.gz
ENV GOROOT_BOOTSTRAP /usr/local/gotemp/go
ADD . /runtime
RUN	cd /usr/local/go/src && patch -p2 < /runtime/go1.6.2-tracenew.patch && ./make.bash
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

WORKDIR $GOPATH
