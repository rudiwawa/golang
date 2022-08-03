FROM golang

COPY main.go /app/main.go

# beri tahu cara runningnya
CMD ["go", "run", "/app/main.go"]