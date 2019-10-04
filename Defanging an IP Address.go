func defangIPaddr(address string) string {
    strlen := len(address)
    for i:=0; i < strlen; i++{
        if address[i] == 46{
            start := string(address[0:i])
            end := string(address[i+1:])
            address = start + "[.]" + end
            strlen += 2
            i+=2
        }
    }
    return address
}
