go build -o pcl ./src

case "$1" in
    repl)
        ./pcl
        ;;
    run)
        ./pcl test.pcl
        ;;
    *)
        echo "brochacho that arg is cooked: $2"
        ;;
esac
