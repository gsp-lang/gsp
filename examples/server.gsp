(ns main
    "/fmt"
    "/net/http")

(def hello (fn [w r]
    (fmt/fprintf w "hello")
    ()))

(def main (fn []
    (http/handle-func "/" hello)
    (http/listen-and-serve ":9090" nil)))
