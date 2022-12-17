FROM ubuntu:focal

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update \
	&& apt-get install -y \
		build-essential \
		software-properties-common \
		ninja-build \
		python3-pip \
		bc \
		wget \
		vim git

# we need meson for libvips build
RUN pip3 install meson

# we use cgif for GIF write
RUN add-apt-repository ppa:lovell/cgif \
	&& apt-get update \
	&& apt-get install -y libcgif-dev

# AVIF&HEIC avif need version > 1.7
RUN add-apt-repository ppa:strukturag/libde265 \
	&& add-apt-repository ppa:strukturag/libheif \
	&& apt-get update \
	&& apt-get install -y libde265-dev \
	&& apt-get install -y libheif-dev

# stuff we need to build our own libvips ... this is a pretty random selection
# of dependencies, you'll want to adjust these
RUN apt-get install -y \
	glib-2.0-dev \
	libexpat-dev \
	librsvg2-dev \
	libpng-dev \
	libjpeg-turbo8-dev \
	libtiff5-dev \
	libexif-dev \
	liblcms2-dev \
	liborc-dev \
	libgsf-1-dev \
	fftw3-dev

# install webp
ARG WEBP_VERSION=1.2.4
RUN wget -P /tmp https://github.com/webmproject/libwebp/archive/refs/tags/v${WEBP_VERSION}.tar.gz
RUN cd /tmp; tar -xvf v${WEBP_VERSION}.tar.gz
RUN apt-get install -y gcc make autoconf automake libtool

RUN cd /tmp/libwebp-${WEBP_VERSION}/; ./autogen.sh; ./configure \
--prefix=/usr/local \
--enable-libwebpmux \
--enable-libwebpdemux \
--enable-libwebpdecoder \
--enable-libwebpextras \
--enable-static

RUN cd /tmp/libwebp-${WEBP_VERSION}/; make; make install

RUN rm "/tmp/v${WEBP_VERSION}.tar.gz"
RUN rm -rf /var/lib/apt/lists/*


WORKDIR /usr/local/src

# build the head of the stable 8.13 branch
ARG VIPS_BRANCH=8.13
ARG VIPS_URL=https://github.com/libvips/libvips/tarball

RUN mkdir libvips-${VIPS_BRANCH} \
        && cd libvips-${VIPS_BRANCH} \
        && wget ${VIPS_URL}/${VIPS_BRANCH} -O - | tar xfz - --strip-components 1

# "--libdir lib" makes it put the library in /usr/local/lib
# we don't need GOI
RUN cd libvips-${VIPS_BRANCH} \
        && rm -rf build \
        && meson build --libdir lib -Dintrospection=false --buildtype release \
        && cd build \
        && ninja \
        && ninja test \
        && ninja install

# pyvips .. handy for testing
RUN pip3 install pyvips

# env vips.so
ENV LD_LIBRARY_PATH $LD_LIBRARY_PATH:/usr/local/lib

# install golang
ARG GO_VERSION=1.16.15
RUN wget -P /tmp "https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz"
RUN tar -C /usr/local -xzf "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"
RUN rm "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR $GOPATH
