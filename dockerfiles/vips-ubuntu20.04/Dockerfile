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
		vim git \
	&& rm -rf /var/lib/apt/lists/*

# we need meson for libvips build
RUN pip3 install meson

# we use cgif for GIF write
RUN add-apt-repository ppa:lovell/cgif \
	&& apt-get update \
	&& apt-get install -y libcgif-dev

# ubuntu20.04 install avif need version > 1.7
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
	libtiff-dev \
	libexif-dev \
	liblcms2-dev \
	liborc-dev

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
RUN wget -P /tmp "https://dl.google.com/go/go1.16.15.linux-amd64.tar.gz"
RUN tar -C /usr/local -xzf "/tmp/go1.16.15.linux-amd64.tar.gz"
RUN rm "/tmp/go1.16.15.linux-amd64.tar.gz"

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR $GOPATH
