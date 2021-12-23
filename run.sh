docker run -d \
        --name learn-go \
        --mount type=bind,source="$(pwd)"/books,target=/books \
        -p 9999:22 \
        learn-go