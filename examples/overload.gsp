(ns main
    "/fmt")

(def print (fn [a] (fmt/println a) ()))

(def print (fn [a b] (fmt/println a b) ()))

(def main (fn []
    (print "foo")
    (print "foo" "bar")))
