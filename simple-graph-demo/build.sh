# go built-in toolchain does not need a Makefile or Maven file
go build


./simple-graph-demo > simple.dot

# Install graphviz on mac
# brew install graphviz

# Visualize the graph
cat simple.dot | dot -Tsvg > output.svg
cat simple.dot | dot -Tpng > output.png

# Cross compile
 GOOS=linux GOARCH=amd64 go build -o simple-graph-demo---amd64-linux
