(ns main
	"/fmt")

(def main (fn []
    (let [[m (/cons 2 (/cons 1 /null))]]
        (fmt/println (/car (/cons 1 /null)))
        (fmt/println (/len (/cons 1 /null)))
        (fmt/println (/car (/map (fn [x] (+ x 20)) m)) (/car m))
        (fmt/println (/len m))
        (fmt/println (/car (/cons 17 m)))
        (fmt/println [1 2 3 4])
        (fmt/println (/car (/cdr (/cdr (/append m (/cdr m))))))
        ())))
