FROM ubuntu
WORKDIR /root/ffmpeg
ADD ./ffmpeg-3.4.2 ./
RUN apt-get update && apt-get  install -y build-essential && apt-get install -y yasm && ./configure && make && make install
WORKDIR /root
RUN rm -rf ffmpeg/