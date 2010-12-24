package dns

// subpackage?

// Convert a string to an resource record
// The string must fit on one line and must be fully formatted
//              IN A 192.168.1.1 // not ok
// miek.nl. 3600 IN A 192.168.1.1 // ok
// miek.nl. IN A 192.168.1.1      // ok, ttl may be omitted
// miek.nl. A 192.168.1.1         // ok, ttl and class omitted
// miek.nl. 3600 A 192.168.1.1    // ok, class omitted
func AtoRR(s string) *RR {
        // up to first whitespace is domainname
        // next word is:
        // <number> -> TTL
        // IN|CH|HS -> Class
        // <rest>   -> Type
        // When the type is seen, we can read the rest
        // of the string in an rr-specific manner
        return nil
}
