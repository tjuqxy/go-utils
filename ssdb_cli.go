package utils

import (
    "fmt"
    "strings"
    "strconv"
)

func SSDBDbsize(addr string) ([]string, error) {
    return callCmd(addr, []byte("6\ndbsize\n\n"))
}

func SSDBSlaveof(addr1, addr2 string) error {
    // 1st addr2 STOP slave
    err := SSDBStopSlave(addr2)
    if err != nil {
        return err
    }

    // 2nd STOP slave
    err  = SSDBStopSlave(addr1)
    if err != nil {
        return err
    }

    // 3rd CHANG_MASTER_TO
    err = SSDBChagneMasterTo(addr1, addr2, "0", "")
    if err != nil {
        return err
    }

    // 4th START slave
    err = SSDBStartSlave(addr1)
    if err != nil {
        return err
    }
    return nil
}

func SSDBAsking(addr string) error {
    output, err := callCmd(addr, []byte("6\nasking\n\n"))
    fmt.Println("Asking", addr, "result:", output)
    return err
}

func SSDBStartSlave(addr string) error {
    _, err := callCmd(addr, []byte("11\nstart_slave\n\n"))
    return err
}

func SSDBStopSlave(addr string) error {
    _, err := callCmd(addr, []byte("10\nstop_slave\n\n"))
    return err
}

func SSDBSlotPremigrating(addr, slot string) error {
    req := "17\nslot_premigrating\n" +
            strconv.Itoa(len(slot)) + "\n" + slot + "\n\n"
    r, err := callCmd(addr, []byte(req))
    fmt.Printf("Addr(%s) mark slot(%s) migrating result: %v\n", addr, slot, r)
    return err
}

func SSDBSlotPreimporting(addr, slot string) error {
    req := "17\nslot_preimporting\n" +
            strconv.Itoa(len(slot)) + "\n" + slot + "\n\n"
    r, err := callCmd(addr, []byte(req))
    fmt.Printf("Addr(%s) mark slot(%s) importing result: %v\n", addr, slot, r)
    return err
}

func SSDBSlotPostmigrating(addr, slot string) error {
    req := "18\nslot_postmigrating\n" +
            strconv.Itoa(len(slot)) + "\n" + slot + "\n\n"
    r, err := callCmd(addr, []byte(req))
    fmt.Printf("Addr(%s) post slot(%s) migrating result: %v\n", addr, slot, r)
    return err
}

func SSDBSlotPostimporting(addr, slot string) error {
    req := "18\nslot_postimporting\n" +
            strconv.Itoa(len(slot)) + "\n" + slot + "\n\n"
    r, err := callCmd(addr, []byte(req))
    fmt.Printf("Addr(%s) post slot(%s) importing result: %v\n", addr, slot, r)
    return err
}

func SSDBMigrateSlot(addr1, addr2, slot string) error {
    addrSplit := strings.Split(addr2, ":")
    if len(addrSplit) != 2 {
        return fmt.Errorf("Addr(%s) is invalid", addr2)
    }
    ip := addrSplit[0]
    port := addrSplit[1]
    req := "12\nmigrate_slot\n" +
            strconv.Itoa(len(slot)) + "\n" + slot + "\n" +
            strconv.Itoa(len(ip)) + "\n" + ip + "\n" +
            strconv.Itoa(len(port)) + "\n" + port + "\n" +
            "1\n0\n" +
            "1\n0\n\n"
    r, err := callCmd(addr1, []byte(req))
    fmt.Printf("Migrate %s to %s result: %v\n", addr1, addr2, r)
    return err
}

func SSDBChagneMasterTo(addr1, addr2, lastSeq, lastKey string) error {
    ip   := ""
    port := ""
    addrSplit := strings.Split(addr2, ":")
    if len(addrSplit) > 1 {
        ip   = addrSplit[0]
        port = addrSplit[1]
    }
    if ip == "" {
        ip = "127.0.0.1"
    }
    if port == "" {
        return fmt.Errorf("Addr(%s) is invalid", addr2)
    }
    req := "16\nchange_master_to\n" +
            strconv.Itoa(len(ip)) + "\n" + ip + "\n" +
            strconv.Itoa(len(port)) + "\n" + port + "\n" +
            strconv.Itoa(len(lastSeq)) + "\n" + lastSeq + "\n" +
            strconv.Itoa(len(lastKey)) + "\n" + lastKey + "\n\n"
    _, err := callCmd(addr1, []byte(req))
    return err
}
