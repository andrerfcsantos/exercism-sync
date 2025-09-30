#!/usr/bin/env gawk -f

BEGIN {
    if ($0 == "") {
        print "One for you, one for me."
    } else {
        print "One for", $0 ", one for me."
    }
}
