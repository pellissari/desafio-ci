FROM scratch

COPY bin/soma /soma

ENTRYPOINT ["/soma","5","5"]