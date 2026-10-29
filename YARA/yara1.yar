rule BLACKOUT-C2-LİTE {
    meta:
        description = "Detects BLACKOUT-C2-LİTE"
        author = "hackpatato"
        severity = "CRITICAL"

    strings:

        $debug1 = "ERROR , no user knowlage" ascii wide
        $debug2 = "data sending is complated" ascii wide
        $func1 = "TcpSend" ascii wide
        $func2 = "Checksandbox" ascii wide
        $package = "AntiSandbox" ascii wide
    condition:
    all of ($debug1, $debug2, $func1, $func2, $package)
}