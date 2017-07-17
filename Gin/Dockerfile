FROM golang
RUN cd  src && mkdir github.com && cd github.com && mkdir gin-gonic && cd gin-gonic && mkdir gin && cd gin \
	&& go get github.com/golang/protobuf/proto \
	&& go get github.com/gin-contrib/sse		\
	&& go get github.com/json-iterator/go \
	&& go get github.com/mattn/go-isatty \
	&& go get github.com/ugorji/go/codec \
	&& go get gopkg.in/go-playground/validator.v8 \
	&& go get gopkg.in/yaml.v2 


ADD ./gin src/github.com/gin-gonic/gin

