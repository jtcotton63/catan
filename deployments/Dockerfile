#
# Build the lambdas using golang
#
FROM golang:1.14.1-stretch
WORKDIR /go/src/app

# Copy go source and other necessary files
# The current directory should now have the 
# following structure:
# app/
#  - pkg/
#  - go.mod
#  - Makefile
COPY ./tmp/pkg/ ./pkg/
COPY ./tmp/go.mod .
COPY ./assets/Makefile .

# Build the lambdas
RUN go mod download
RUN make build



#
# Deploy the lambdas to the local stage (the localstack AWS mock)
#
FROM node:13.12.0-stretch

# Note: the node user and group are created by the node image.
# RUN chown -R node:node /home/node
# USER node:node
WORKDIR /work

# Copy the go files built in the previous stage
# as well as other files that will be needed
COPY --from=0 /go/src/app/bin ./bin/
COPY ./assets/deploy.sh .
COPY ./assets/serverless.yml .
RUN chmod +x ./deploy.sh

# Install serverless
# Note that npm is just used as a quick means
# to install serverless
RUN npm init -y
RUN npm install -g serverless
RUN npm install --save-dev serverless-localstack

# Clone wait-for-it
RUN git clone https://github.com/vishnubob/wait-for-it.git
RUN chmod +x ./wait-for-it/wait-for-it.sh

# Exec the command
CMD ./deploy.sh